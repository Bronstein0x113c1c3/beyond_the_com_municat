package trackingserv

import (
	"fmt"
	"tracking_server/tracking_protoc"
	"tracking_server/typing"

	"github.com/reactivego/multicast"
	// "github.com/SierraSoftworks/multicast"
	// "github.com/SierraSoftworks/multicast"
	// "github.com/sierrasoftworks/multicast"
)

type TrackerServer struct {
	addr    string
	port    int
	Tracker *multicast.Chan
	tracking_protoc.UnimplementedTrackingServer
}

func Init(addr string, port int, max_buffer, max_users int) *TrackerServer {
	return &TrackerServer{
		addr:    addr,
		port:    port,
		Tracker: multicast.NewChan(max_buffer, max_users), //max would be 20 client
	}
}
func (t *TrackerServer) String() string {
	return fmt.Sprintf("%v:%v", t.addr, t.port)
}
func (t *TrackerServer) Notify(n *typing.Notification) {
	t.Tracker.Send(n)
}
func (t *TrackerServer) Close() error {
	var err error
	t.Tracker.Close(err)
	return err
}
