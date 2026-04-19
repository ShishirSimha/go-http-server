package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//Use UDP instead of TCP
	listener, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	//Dial the UDP connection
	conn, err := net.DialUDP("udp", nil, listener)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading from stdin:", err)
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Println("error writing to udp:", err)
		}
	}
}