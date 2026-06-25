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
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("cant create socket")

	}

	for {
		conn, err := listner.Accept()

		if err != nil {
			fmt.Println("handshake unsuccessfulll")
		}

		go handleConn(conn)
	}
}

func sendResponse(conn net.Conn, status string, body string) {
	getMessage := fmt.Sprintf("HTTP/1.1 %s\r\n"+"Content-Type: text/html\r\n"+"Content-Length: %v\r\n\r\n%s", status, len(body), body)
	conn.Write([]byte(getMessage))

}

func handleGet(words []string, conn net.Conn) {
	switch words[1] {
	case "/":
		{
			sendResponse(conn, "200 OK", `<h1>Bhopdikeee</h1> 
			<form method="POST" action="/submit">
			<input name="username">
			<button type="submit">send</button>	
			</form>`)
		}
	case "/about":
		{
			sendResponse(conn, "200 OK", "<h1>about</h1>")
		}
	default:
		{
			sendResponse(conn, "404 Not Found", "<h1>404</h1>")
		}
	}

}

func handlePost(reader *bufio.Reader, conn net.Conn) {
	contentLength := 0
	// var err error
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error")
		}

		if line == "\r\n" {
			body := make([]byte, contentLength)
			_, err = reader.Read(body)
			if err != nil {
				fmt.Println("error reading body")
			}
			fmt.Println(string(body))

			parts := strings.Split(string(body), "=")
			if len(parts) != 2 {
				sendResponse(conn, "400 Bad Request", "<h1>Bad Request</h1>")
				return
			}

			userName := parts[1]
			respo := fmt.Sprintf("<h1>Hello %s</h1>", userName)
			sendResponse(conn, "200 OK", respo)
			break
		}

		if strings.HasPrefix(line, "Content-Length") {
			fmt.Sscanf(line, "Content-Length: %d", &contentLength)
			fmt.Println(contentLength)
		}
	}

}

func handleConn(conn net.Conn) {
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

	fmt.Printf("method = %v\npath = %v\nversion = %v\n", words[0], words[1], words[2])

	switch words[0] {

	case "GET":
		{
			handleGet(words, conn)
		}

	case "POST":
		{
			handlePost(reader, conn)
		}
	}

}

func main() {
	clientHandler()
}
