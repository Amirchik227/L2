package telnet

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Client struct {
	Conn net.Conn
}

// NewTelnetClient инициализирует новый клиент с определенным порт, хостом и таймаутом соединения
func NewTelnetClient(host string, port string, timeout int) (*Client, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	err = conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return nil, err
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected")

	return &Client{Conn: conn}, nil
}

// Start запускает tcp клиент и запускает свой обработчик
func (c *Client) Start() {
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go handle(c.Conn, errChan)

	for {
		select {
		case <-sigChan:
			fmt.Println("\n Закрываем соединение")
			return
		case err := <-errChan:
			log.Printf("%s\n", err)
			fmt.Fprintf(c.Conn, "STOP")

			return
		}
	}
}

// handle является простейшим обработчиком клиента, представляющий собой строку с приглашением для ввода
func handle(conn net.Conn, errChan chan error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			errChan <- err
			return
		}

		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			errChan <- err
			return
		}

		fmt.Print("->: " + message)
	}
}
