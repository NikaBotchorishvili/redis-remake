package client

import (
	"fmt"
	"net"
	"strings"

	"github.com/NikaBotchorishvili/redis-remake/internal"
)


func HandleMessage(conn net.Conn, store *internal.Store, message string) (net.Conn, error) {

	slicedMessage := strings.Fields(message)
	var err error
	cmd, err := parseCommand(slicedMessage)
	switch cmd {
	case SET_COMMAND:
		err = handleSET(slicedMessage, store)
	case GET_COMMAND:
		var val string

		val, err = handleGET(slicedMessage, store)
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error: %s\n", err.Error())))
		} else {
			conn.Write([]byte(fmt.Sprintf("Value: %s\n", val)))
		}
	case DELETE_COMMAND:
		if len(slicedMessage) != 2 {
			err = fmt.Errorf("expected 2 params, got %d", len(slicedMessage))
			conn.Write([]byte("Error: Expected format DELETE key\n"))
			return conn, err
		}
		key := slicedMessage[1]
		store.Delete(key)
		conn.Write([]byte("OK\n"))
		fmt.Println("DELETE command")

	default:
		err = fmt.Errorf("unknown command %v", cmd)
		conn.Write([]byte(fmt.Sprintf("Unknown command: %s\n", cmd)))
		fmt.Println("Unknown command")
	}

	return conn, err
}


func handleGET(slicedMessage []string, store *internal.Store) (string, error){
	var err error
	if len(slicedMessage) != 2 {
		err = fmt.Errorf("expected 2 params, got %d", len(slicedMessage))
		return "", err
	}
	key := slicedMessage[1]
	val, err := store.Get(key)
	fmt.Println()

	if err != nil {
		return "", err
	}
	return val, nil
}

func handleSET(slicedMessage []string, store *internal.Store) error{
	if len(slicedMessage) != 3 {
		err := fmt.Errorf("expected 3 params, got %d", len(slicedMessage))
		return err
	}
	key, value := slicedMessage[1], slicedMessage[2]
	store.Set(key, value)
	fmt.Printf("SET key: %s to value %s\n", key, value)

	return nil
}

func handleDELETE(){

}

func parseCommand(slicedMessage []string) (string, error) {

	if len(slicedMessage) == 0 {
		err := fmt.Errorf("empty command")
		return "", err
	}

	cmd := slicedMessage[0]
	if len(cmd) == 0 {
		err := fmt.Errorf("invalid command %v", cmd)
		return "", err
	}


	return cmd, nil
} 