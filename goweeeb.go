package main

import (
	"bufio"
	"fmt"
	// "io"
	// "strconv"

	// "io"
	"net"
	// "os"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string

	Headers map[string]string
	Body    []byte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func clientHandler() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("cant create socket")

	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("handshake unsuccessfulll")
		}

		go handleConn(conn)
	}
}

func sendResponse(conn net.Conn, status string, conType string, body string) {
	getMessage := fmt.Sprintf("HTTP/1.1 %s\r\n"+"Content-Type: %s\r\n"+"Content-Length: %v\r\n\r\n%s", status, conType, len(body), body)
	conn.Write([]byte(getMessage))

}

func handleConn(conn net.Conn) {
	req := Request{
		Headers: make(map[string]string),
	}

	fmt.Printf("client connected :%v\n", conn.RemoteAddr())
	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("some happend")
	}

	words := strings.Fields(line)

	if len(words) < 3 {
		fmt.Println("invalid request")
		return
	}

	req.Method, req.Path, req.Version = words[0], words[1], words[2]

	fmt.Printf("method = %v\npath = %v\nversion = %v\n", req.Method, req.Path, req.Version)

	switch req.Method {

	case "GET":
		{
			handleGet(&req, conn)
		}

	case "POST":
		{
			handlePost(&req, reader, conn)
		}
	}

}

func main() {
	clientHandler()
}
