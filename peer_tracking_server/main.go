package main

import (
	"fmt"
	"net"
	"tracking_server/interceptor"
	"tracking_server/tracking_protoc"
	trackingserv "tracking_server/tracking_serv"

	"google.golang.org/grpc"
)

func main() {

	serv := trackingserv.Init("", 8080, 512, 20)
	grpc_helper := grpc.NewServer(grpc.ChainStreamInterceptor(interceptor.Limiter(serv), interceptor.Welcome(serv)))
	tracking_protoc.RegisterTrackingServer(grpc_helper, serv)
	lis, err := net.Listen("tcp", fmt.Sprintf("%v", serv))
	if err != nil {
		return
	}
	// go func() {

	// }()
	grpc_helper.Serve(lis)
	// grpc_helper.Serve()
}
