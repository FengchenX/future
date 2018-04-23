
package main

import (
	"fmt"
	"bufio"
	"log"
	"net"

)
func main() {
	ln, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn,err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnect(conn)
	}
}


func handleConnect(conn net.Conn) {
	//todo 
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		conn.Write([]byte("hello,world\n"))
	}
}