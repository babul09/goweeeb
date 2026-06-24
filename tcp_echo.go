package main

import (
	// "bufio"
	"fmt"
	"io"
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

func handleConn(conn net.Conn) {
	fmt.Printf("client connected :%v\n", conn.RemoteAddr())
	defer conn.Close()

	data, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("some happend")
	}

	line := string(data)

	line = strings.Split(line, "\r\n")[0]

	words := strings.Fields(line)

	if len(words) < 3 {
		fmt.Println("invalid request")
		return
	}

	fmt.Printf("method = %v\npath = %v\nversion = %v\n", words[0], words[1], words[2])

}

func main() {
	clientHandler()
}
