package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":4455")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "What's your name?\n> ")

	r := bufio.NewReader(conn)

	name, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "Hello, %s\n", name)
}
