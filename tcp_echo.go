package main

import (
	"bufio"
	"fmt"
	// "io"
	"net"
	"strings"
	// "os"
)

func clientHandler() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("cant crete socket")

	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("handshake unsuccessfulll")
		}

		go handleConn(conn)
	}
}

func SendRequest(conn net.Conn, status string, body string) {
	getMessage := fmt.Sprintf("HTTP/1.1 %s\r\n"+"Content-Type: text/html\r\n"+"Content-Length: %v\r\n\r\n%s", status, len(body), body)
	conn.Write([]byte(getMessage))

}

func handleConn(conn net.Conn) {
	fmt.Printf("client connected :%v\n", conn.RemoteAddr())
	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("some happend")
	}

	// line = strings.Split(line, "\r\n")[0]

	words := strings.Fields(line)

	if len(words) < 3 {
		fmt.Println("invalid request")
		return
	}

	fmt.Printf("method = %v\npath = %v\nversion = %v\n", words[0], words[1], words[2])

	if words[0] == "GET" {

		switch words[1] {
		case "/":
			{
				SendRequest(conn, "200 OK", "<h1>hello<h1>")
			}
		case "/about":
			{
				getMessage := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: 14\r\n\r\n<h1>About</h1>")

				conn.Write(getMessage)

			}
		default:
			{
				getMessage := []byte("HTTP/1.1 404 Not Found\r\nContent-Type: text/html\r\nContent-Length: 12\r\n\r\n<h1>404</h1>")

				conn.Write(getMessage)

			}
		}

	}

}

func main() {
	clientHandler()
}
