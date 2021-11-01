package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main3() {
	listen, err := net.Listen("tcp", "10.168.8.86:8080")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for  {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue;
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("catch err", err)
		}

	}()
	reader := bufio.NewReader(conn)
	var buff [1024]byte
	fmt.Printf("%s:%p\n", "buff", &buff)
	for  {
		sl := buff[:]
		fmt.Printf("%s:%p\n", "slice", sl)
		n, err := reader.Read(sl);
		if err == io.EOF{
			break
		}
		if err != nil {
			panic(err)
		}
		recvstr := string(buff[:n])
		fmt.Println("client string:", recvstr)
		conn.Write([]byte(recvstr))
	}
}