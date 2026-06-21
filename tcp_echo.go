package main

import (
	// "bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func createServer() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("oh shit, not good, server ded")
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("batloll server listeing on post 8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("me go die now")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Client Connected %v\n", conn.RemoteAddr().String())

	_, err := io.Copy(conn, conn)
	if err != nil {
		fmt.Println("cant write back")
	}
	fmt.Printf("client dissssssconected %s\n", conn.RemoteAddr().String())

}

func main() {
	//conn, err := net.Dial("tcp", "golang.org:80")
	//if err != nil {
	//	fmt.Println("could not create connection")
	//}
	//fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")

	//status, err := bufio.NewReader(conn).ReadString('\n')

	//fmt.Println(status)

	createServer()

}
