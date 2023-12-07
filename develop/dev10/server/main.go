package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func main() {
	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Server is working")
	defer li.Close()
	var wg sync.WaitGroup

	conn, err := li.Accept()
	if err != nil {
		log.Println(err)
	}

	wg.Add(1)
	go handle(conn, &wg)
	wg.Wait()

}

// handle представляет собой простейший обработчик tcp соединения, который выводит данные, записанные клиентом в сокет
func handle(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()
	notify := make(chan error)
	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				notify <- err
				return
			}
			fmt.Println(data)
			fmt.Fprintf(conn, "Server read massage: %s", data)
		}
	}()

	for {
		select {
		case err := <-notify:
			if io.EOF == err {
				fmt.Println("connection dropped", err)
				return
			}
		}
	}
}
