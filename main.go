package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type message struct {
	user string
	body string
}

var cs = [](chan message){}

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
	fmt.Fprintf(conn, "bot> What's your name?: ")

	r := bufio.NewReader(conn)

	user, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	user = strings.TrimSpace(user)
	fmt.Fprintf(conn, "bot> Hello, %s!\n", user)

	c := make(chan message)
	cs = append(cs, c)

	go func() {
		for {
			body, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			body = strings.TrimSpace(body)
			if len(body) == 0 {
				continue
			}
			broadcast(user, body)
		}
	}()

	for {
		m := <-c
		fmt.Fprintf(conn, "%s> %s\n", m.user, m.body)
	}
}

func broadcast(user string, body string) {
	m := message{
		user: user,
		body: body,
	}

	for _, c := range cs {
		c <- m
	}
}
