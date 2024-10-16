package interceptor

import (
	"context"
	"log"
	trackingserv "tracking_server/tracking_serv"
	"tracking_server/typing"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type wrapped struct {
	grpc.ServerStream
	id string
}

func (w *wrapped) Context() context.Context {
	return context.WithValue(w.ServerStream.Context(), typing.T("channel"), w.id)
}

// func Passcode
func Welcome(s *trackingserv.TrackerServer) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		id := uuid.NewV4()
		log.Printf("slot %v is opened!!!", id)
		handler(srv, &wrapped{ss, id.String()})
		return nil

	}
}

func Limiter(s *trackingserv.TrackerServer) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		p, ok := peer.FromContext(ss.Context())
		log.Printf("%v is called... \n", info.FullMethod)
		if !ok {
			log.Println("Cannot get any context.....")
		}
		log.Printf("%v wants to connect....: ", p.Addr.String())

		if s.Check() {
			// log.Printf("remaining %v connections... \n", a)
			log.Println("Out of slots")
			return status.Error(codes.ResourceExhausted, "out of slots, cancelled!!!")
		}
		log.Println("Accepted!!!")
		log.Printf("Remaining %v slots \n", s.Remaining())
		s.GetIn()
		handler(srv, ss)
		return nil
	}

}
