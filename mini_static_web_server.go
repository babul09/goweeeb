package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	// "io"
	"net"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reader(path string) (readbuf []byte) {

	file, err := os.Open(path)

	check(err)
	var allbyte []byte
	defer file.Close()
	r := bufio.NewReader(file)
	buff := make([]byte, 1024)
	for {
		n, err := r.Read(buff)
		if n > 0 {
			allbyte = append(allbyte, buff[:n]...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(n)
		}
	}
	return allbyte

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

func handleGet(words []string, conn net.Conn) {
	switch words[1] {
	case "/":
		{
			sendResponse(conn, "200 OK", "text/html", string(reader("./public/public/index.html")))
		}
	case "/styles.css":
		sendResponse(
			conn,
			"200 OK",
			"text/css",
			string(reader("./public/public/styles.css")),
		)
	case "/script.js":
		sendResponse(
			conn,
			"200 OK",
			"application/javascript",
			string(reader("./public/public/script.js")),
		)
	case "/about":
		{
			sendResponse(conn, "200 OK", "text/html", "<h1>about</h1>")
		}
	default:
		{
			sendResponse(conn, "404 Not Found", "text/html", "<h1>404</h1>")
		}
	}

}

func handlePost(reader *bufio.Reader, conn net.Conn) {
	contentLength := 0
	headers := make(map[string]string)
	var err error
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error")
		}

		if line == "\r\n" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		parts[1] = strings.TrimSpace(parts[1])

		headers[parts[0]] = parts[1]

	}

	contentLength, err = strconv.Atoi(headers["Content-Length"])
	if err != nil {
		fmt.Println("error in length conv")
		return
	}

	fmt.Println(contentLength)
	body := make([]byte, contentLength)
	_, err = reader.Read(body)
	if err != nil {
		fmt.Println("error reading body")
	}
	fmt.Println(string(body))

	parts := strings.SplitN(string(body), "=", 2)
	if len(parts) != 2 {
		sendResponse(conn, "400 Bad Request", "text/html", "<h1>Bad Request</h1>")
		return
	}

	userName := parts[1]
	respo := fmt.Sprintf("<h1>Hello %s</h1>", userName)
	sendResponse(conn, "200 OK", "text/html", respo)

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
