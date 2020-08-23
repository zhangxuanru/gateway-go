package main

import (
	"log"
	"net"
	"gateway/pack"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		log.Fatal("net.Dial error:", err)
	}
	pack.Encode(conn, "hello world!!!")
}
