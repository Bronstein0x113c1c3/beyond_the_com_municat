package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"tracking_server/interceptor"
	"tracking_server/tracking_protoc"
	trackingserv "tracking_server/tracking_serv"

	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	serv := trackingserv.Init("", 8080, 512, 20)
	grpc_helper := grpc.NewServer(grpc.ChainStreamInterceptor(interceptor.Limiter(serv), interceptor.Welcome(serv)))
	tracking_protoc.RegisterTrackingServer(grpc_helper, serv)
	lis, err := net.Listen("tcp", fmt.Sprintf("%v", serv))
	if err != nil {
		return
	}
	// go func() {

	// }()
	go func() {
		grpc_helper.Serve(lis)
	}()
	<-ctx.Done()
	serv.Close()
	grpc_helper.GracefulStop()
	serv.Clear()
	log.Println("stopped gracefully")
	// grpc_helper.Serve()
}
