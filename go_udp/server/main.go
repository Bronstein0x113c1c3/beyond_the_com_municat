package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
)

func connection_spawning(station *net.UDPConn, ctx context.Context) chan *net.UDPAddr {
	// net.Pipe()
	// station.SetDeadline()
	channel := make(chan *net.UDPAddr, 10)
	go func() {
		defer log.Println("channel is closed!!!!")
		defer close(channel)
		// defer station.Close()
		for {
			data := make([]byte, 1)
			n, conn, err := station.ReadFromUDP(data)
			if err != nil {
				log.Println("stopped receiving")
				return
			}
			// connection just spawning...., get some request....
			if n <= 0 {
				log.Println("a new connection is spawned!!!!")
				channel <- conn
			}
		}
	}()
	go func() {
		<-ctx.Done()
		log.Println("received signal from keyboard!!!!")
		station.Close()
	}()
	return channel
}

func connection_receiving(station *net.UDPConn, ctx context.Context) (chan *MSG, chan *Hello, chan *Goodbye) {
	msg_chan := make(chan *MSG)
	hello_chan := make(chan *Hello)
	goodbye_chan := make(chan *Goodbye)
	go func() {

	}()
	go func() {
		<-ctx.Done()
		log.Println("received signal from keyboard!!!!")
		station.Close()
	}()
	return msg_chan, hello_chan, goodbye_chan
}

// func send_packets(station *net.UDPConn)
func main() {
	// net.ListenPacket()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	station, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("initiated")
	spawner := connection_spawning(station, ctx)
	for conn := range spawner {
		station.WriteToUDP([]byte("Hello!!!!"), conn)

	}
}

/*
	handshaking...

*/
