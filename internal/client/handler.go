package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/NikaBotchorishvili/redis-remake/internal"
)
func HandleClient (conn net.Conn, store *internal.Store) {
	fmt.Println("New connection from: ", conn.RemoteAddr().String())
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for{
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		message = strings.TrimSpace(message)

		HandleMessage(conn, store,  message)

	}
}