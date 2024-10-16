package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"p2p_client/tracking_protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	//connect to server

	addr := "abc"
	port := 1234
	conn, err := grpc.NewClient(":8080", grpcOpts...)
	if err != nil {
		return
	}
	client, err := tracking_protoc.NewTrackingClient(conn).Add(ctx, &tracking_protoc.IncomingRequest{
		Ip:   addr,
		Port: int32(port),
	})
	if err != nil {
		return
	}
	info, err := client.Recv()
	if err != nil {
		return
	}
	log.Println(info)
	go func() {
		for {
			info, err := client.Recv()
			if err != nil {
				return
			}
			if !info.Status {
				log.Printf("got out %v \n", info)
				continue
			} else {
				log.Println(info)
			}

		}
	}()
	<-ctx.Done()
}
