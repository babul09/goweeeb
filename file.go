package main

import (
	"bufio"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
)

func readFile(path string) (readbuf []byte, mtype string, err error) {

	file, err := os.Open("public/public" + path)

	if err != nil {
		fmt.Println("invalid path")
		return
	}
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
			return nil, "", err
		}
	}

	ext := filepath.Ext("public/public" + path)

	mimeType := mime.TypeByExtension(ext)

	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	return allbyte, mimeType, err

}
