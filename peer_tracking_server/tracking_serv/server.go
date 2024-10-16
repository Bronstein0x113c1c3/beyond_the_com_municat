package trackingserv

import (
	"fmt"
	"sync"
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
	limiter chan struct{}
	tracking_protoc.UnimplementedTrackingServer
	max_user       int
	receipent_list *sync.Map
}

func Init(addr string, port int, max_buffer, max_users int) *TrackerServer {
	return &TrackerServer{
		addr:           addr,
		port:           port,
		Tracker:        multicast.NewChan(max_buffer, max_users), //max would be 20 client
		limiter:        make(chan struct{}, max_users),
		max_user:       max_users,
		receipent_list: &sync.Map{},
	}
}
func (t *TrackerServer) Remaining() int {
	return t.max_user - len(t.limiter)
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
func (t *TrackerServer) Check() bool {
	return len(t.limiter) > t.max_user
}

func (t *TrackerServer) add_someone(id string, info typing.Notification) {
	t.receipent_list.LoadOrStore(id, info)
}

func (t *TrackerServer) leave(id string) {
	t.receipent_list.Delete(id)
	<-t.limiter
}

func (t *TrackerServer) GetIn() {
	t.limiter <- struct{}{}
}
func (t *TrackerServer) list() []*tracking_protoc.IncomingRequest {
	var res []*tracking_protoc.IncomingRequest
	t.receipent_list.Range(func(key, value any) bool {
		v := value.(typing.Notification)
		res = append(res, &tracking_protoc.IncomingRequest{
			Ip:   v.Addr,
			Port: int32(v.Port),
		})
		return true
	})
	return res
}
