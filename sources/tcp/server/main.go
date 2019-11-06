package main

import (
	"bufio"
	"flag"
	"log"
	"net"
)



func handle (c net.Conn) {
	defer c.Close()
	log.Println("acctepted connection")

	reader  := bufio.NewReader(c)
	for {
		// Waiting for ping
		line, _, _ := reader.ReadLine()

		if string(line) != ""{
			log.Println("recieved from client", string(line))
		}
	}
}


func health(){
	ln, _ := net.Listen("tcp", ":4343")

	for {
		_, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port","4242","")
	flag.Parse()

	log.Println("starting tcp listener")

	go health()

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil{
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}