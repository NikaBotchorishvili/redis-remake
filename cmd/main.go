package main

import (
	"fmt"
	"log"
	"net"

	"github.com/NikaBotchorishvili/redis-remake/internal"
	"github.com/NikaBotchorishvili/redis-remake/internal/client"
)


const DEFAULT_PORT = ":5825"

func main () {

	ln, err := net.Listen("tcp", DEFAULT_PORT)
	
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ln.Close()
	srv := &internal.Server{}
	store := internal.NewStore()
	srv.Listener = ln

	fmt.Printf("Server is running on port: %s\n", DEFAULT_PORT)

	for {
		conn, err := srv.Listener.Accept()

		if err != nil {
			log.Println(err.Error())
			continue
		}
		go client.HandleClient(conn, store)
		
	}
}

