package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8081,
	})
	if err != nil {
		log.Fatal("DialUdp error:", err)
	}
	for i := 1; i <= 100; i++ {
		_, err := conn.Write([]byte("hello world!!!"))
		if err != nil {
			log.Fatal("write error:", err)
		}
		r := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(r)
		if err != nil {
			log.Fatal("read error:", err)
		}
		fmt.Println("recv string :", string(r[:n]))
	}

}
