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

func SendResponse(conn net.Conn, status string, body string) {
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

	words := strings.Fields(line)

	if len(words) < 3 {
		fmt.Println("invalid request")
		return
	}

	ContentLength := 0

	fmt.Printf("method = %v\npath = %v\nversion = %v\n", words[0], words[1], words[2])

	if words[0] == "GET" {

		switch words[1] {
		case "/":
			{
				SendResponse(conn, "200 OK", `<h1>Bhopdikeee</h1> 
			<form method="POST" action="/submit">
			<input name="username">
			<button type="submit">send</button>	
			</form>`)
			}
		case "/about":
			{
				SendResponse(conn, "200 OK", "<h1>about</h1>")
			}
		default:
			{
				SendResponse(conn, "404 Not Found", "<h1>404</h1>")
			}
		}

	} else if words[0] == "POST" {
		for {
			line, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("read error")
			}

			if line == "\r\n" {
				body := make([]byte, ContentLength)
				_, err = reader.Read(body)
				if err != nil {
					fmt.Println("error reading body")
				}
				fmt.Println(string(body))

				parts := strings.Split(string(body), "=")
				userName := parts[1]
				respo := fmt.Sprintf("<h1>Hello %s</h1>", userName)
				SendResponse(conn, "200 OK", respo)
				break
			}

			if strings.HasPrefix(line, "Content-Length") {
				fmt.Sscanf(line, "Content-Length: %d", &ContentLength)
				fmt.Println(ContentLength)
			}
		}

	}

}

func main() {
	clientHandler()
}
