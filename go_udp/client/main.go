package main

import (
	"fmt"
	"net"
	"os"
)

// type packet interface {
// 	marshal() []byte
// 	type_of_packet() string
// }

// type conn_test struct {
// 	id string

// }

func main() {
	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// conn_testing(conn)
	data := make([]byte, 100)
	n, _ := conn.Read(data)
	fmt.Println(data[:n])
}
