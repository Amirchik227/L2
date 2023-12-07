package main

import (
	"flag"
	"log"

	"dev10/telnet"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	host := flag.String("host", "localhost", "хост")
	port := flag.String("port", "3000", "порт")
	timeout := flag.Int("timeout", 10, "таймаут на подключение к серверу")

	flag.Parse()

	client, err := telnet.NewTelnetClient(*host, *port, *timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Conn.Close()

	client.Start()
}
