package main

import (
	"bufio"
	"io"
	"os"
)

func readFile(path string) (readbuf []byte) {

	file, err := os.Open("public/public" + path)

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
