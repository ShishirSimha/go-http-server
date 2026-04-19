package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		var currentLine string
		//Read the connection 8 bytes at a time
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)

			if n > 0 {
				parts := strings.Split(string(data[:n]), "\n")
				for i := 0; i < len(parts)-1; i++ {
					line := currentLine + parts[i]
					line = strings.TrimSuffix(line, "\r")
					out <- line
					currentLine = ""
				}
				currentLine += parts[len(parts)-1]
			}

			if err != nil { //on error (like io.EOF), gracefully exit
				break
			}
		}

		if currentLine != "" {
			currentLine = strings.TrimSuffix(currentLine, "\r")
			out <- currentLine
		}

	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println("connection accepted")

		//Uses our getLinesChannel function to read lines from the connection.
		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Printf("read %s\n", line)
		}

		fmt.Println("connection closed")
	}
}
