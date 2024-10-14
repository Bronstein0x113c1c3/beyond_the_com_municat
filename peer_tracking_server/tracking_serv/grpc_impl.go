package trackingserv

import (
	"context"
	"tracking_server/tracking_protoc"

	"github.com/reactivego/multicast"
)

func (t *TrackerServer) Add(*tracking_protoc.IncomingRequest, tracking_protoc.Tracking_AddServer) error {
	t.Tracker.NewEndpoint(multicast.ReplayAll)
}
func (t *TrackerServer) Out(context.Context, *tracking_protoc.StopRequest) (*tracking_protoc.Notify, error) {

}
