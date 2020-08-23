package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8081,
	})
	if err != nil {
		log.Fatal("listenudp error:", err)
	}
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			log.Println("readFromUDP error:", err)
			break
		}
		go func() {
			_, err := listen.WriteToUDP([]byte("recv success"), addr)
			fmt.Println("recv string:", string(buf[:n]))
			if err != nil {
				log.Println("WriteToUDP error :", err)
			}
		}()
	}
}
