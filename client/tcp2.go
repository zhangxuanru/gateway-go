package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:9000")
	defer conn.Close()
	if e != nil {
		log.Fatal("dial error:", e)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		s, e := reader.ReadString('\n')
		if e != nil {
			log.Fatal("readString err:", e)
		}
		s = strings.TrimSpace(s)
		if s == "q" {
			break
		}
		_, e = conn.Write([]byte(s))
		if e != nil {
			log.Fatal("write error:", e)
		}
	}
}
