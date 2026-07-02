package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func handleGet(req *Request, conn net.Conn) {
	switch req.Path {
	case "/":
		{
			sendResponse(conn, "200 OK", req.Headers["Content-Type"], string(readFile("/index.html")))
		}
	case "/styles.css":
		sendResponse(
			conn,
			"200 OK",
			req.Headers["Content-Type"],
			string(readFile(req.Path)),
		)
	case "/script.js":
		sendResponse(
			conn,
			"200 OK",
			"application/javascript",
			string(readFile(req.Path)),
		)
	case "/about":
		{
			sendResponse(conn, "200 OK", "text/html", "<h1>about</h1>")
		}
	default:
		{
			sendResponse(conn, "404 Not Found", "text/html", string(readFile(req.Path)))
		}
	}

}

func handlePost(req *Request, reader *bufio.Reader, conn net.Conn) {

	contentLength := 0
	// headers := make(map[string]string)
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

		req.Headers[parts[0]] = parts[1]

	}

	contentLength, err = strconv.Atoi(req.Headers["Content-Length"])
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
