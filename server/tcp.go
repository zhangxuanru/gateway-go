package main

import (
	"fmt"
	"gateway/pack"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("net.Listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go Proc(conn)
	}
}

func Proc(conn net.Conn) {
	defer conn.Close()
	body, err := pack.Decode(conn)
	if err!=nil{
		log.Fatal("pack decode error:",err)
	}
	fmt.Println("recv string:",string(body))
}

