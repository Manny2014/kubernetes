package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)





func main() {
	_ = context.Background()

	var server string

	flag.StringVar(&server, "server", "0.0.0.0:4242","")
	flag.Parse()
	log.Println("connecting client to ", server)
	conn, err := net.Dial("tcp", server)
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}


	for {
		time.Sleep(20000)
		// Sent server ping
		fmt.Fprintf(conn, "ping\n")
	}
}