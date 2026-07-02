package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func serveFile(req *Request, conn net.Conn) {
	path := req.Path

	if path == "/" {
		path = "/index.html"
	}

	data, mime, err := readFile(req.Path)
	if err != nil {
		fmt.Println(" file not found ")
		return
	}
	sendResponse(conn, "200 OK", mime, string(data))
}

func handleGet(req *Request, conn net.Conn) {
	serveFile(req, conn)
}

func readRequest(conn net.Conn) (*Request, error) {
	req := &Request{
		Headers: make(map[string]string),
	}

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("some happend")
	}

	words := strings.Fields(line)

	if len(words) < 3 {
		fmt.Println("invalid request")
		return nil, err
	}

	req.Method, req.Path, req.Version = words[0], words[1], words[2]

	fmt.Printf("method = %v\npath = %v\nversion = %v\n", req.Method, req.Path, req.Version)

	// headers := make(map[string]string)
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

	contentLengthStr, ok := req.Headers["Content-Length"]
	if ok {
		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			fmt.Println("error in length conv")
			return nil, err
		}

		req.Body = make([]byte, contentLength)
		_, err = reader.Read(req.Body)
		if err != nil {
			fmt.Println("error reading body")
		}
	}
	return req, nil
}

func handlePost(req *Request, conn net.Conn) {

	fmt.Println(string(req.Body))

	parts := strings.SplitN(string(req.Body), "=", 2)
	if len(parts) != 2 {
		sendResponse(conn, "400 Bad Request", "text/html", "<h1>Bad Request</h1>")
		return
	}

	userName := parts[1]
	respo := fmt.Sprintf("<h1>Hello %s</h1>", userName)
	sendResponse(conn, "200 OK", "text/html", respo)

}
