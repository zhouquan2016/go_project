package main

import (
	"fmt"
	"net"
)

func main4() {
	conn, err := net.Dial("tcp", "10.168.8.86:8080")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
	}
}