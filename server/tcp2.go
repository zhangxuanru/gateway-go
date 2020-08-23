package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:9000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Fatal("accept:", e)
		}
		go Procc(conn)
	}

}

func Procc(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatal("read error:", err)
			break
		}
		fmt.Println("recv string:", string(buf[:n]))
	}
}
