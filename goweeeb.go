package main

import (
	"fmt"

	"net"
)

type Request struct {
	Method  string
	Path    string
	Version string

	Headers map[string]string
	Body    []byte
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
	_, err := conn.Write([]byte(getMessage))

	if err != nil {
		fmt.Println("could not send response")
		return
	}

}

func handleConn(conn net.Conn) {

	fmt.Printf("client connected :%v\n", conn.RemoteAddr())
	defer conn.Close()

	req, err := readRequest(conn)
	if err != nil {
		return
	}

	switch req.Method {

	case "GET":
		{
			handleGet(req, conn)
		}

	case "POST":
		{
			handlePost(req, conn)
		}
	}

}

func main() {
	clientHandler()
}
