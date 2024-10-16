package trackingserv

import (
	"log"
	"tracking_server/tracking_protoc"
	"tracking_server/typing"
)

type foreach func(value interface{}, err error, closed bool) bool

/*
plan:

 1. receive information from the client, use interceptor or request

 2. send other receipent list to the client

 3. add client to receipent list then notify again

 4. hold the connection to the client
*/
func (t *TrackerServer) Add(r *tracking_protoc.IncomingRequest, s tracking_protoc.Tracking_AddServer) error {
	ID := first(r, s, t)
	new_chan, err := t.Tracker.NewEndpoint(0)
	if err != nil {
		return err
	}
	go func() {
		<-s.Context().Done()
		log.Printf("context called for %v \n", ID)
		new_chan.Cancel()
	}()

	for_each := func(s tracking_protoc.Tracking_AddServer) foreach {
		return func(value interface{}, err error, closed bool) bool {
			if closed {
				return false
			}
			n := value.(*typing.Notification)
			m := &tracking_protoc.Notify{
				Status: n.Status,
				Request: &tracking_protoc.Notify_Other{
					Other: &tracking_protoc.OtherNotify{
						Id: n.ID,
						Info: &tracking_protoc.IncomingRequest{
							Ip:   n.Addr,
							Port: int32(n.Port),
						},
					},
				},
			}
			// if.Send()
			if errx := s.Send(m); errx != nil {
				return false
			}
			return true
		}

	}
	new_chan.Range(for_each(s), 0)
	defer func() {
		t.leave(ID)
		t.Tracker.Send(&typing.Notification{
			Addr:   r.Ip,
			Port:   int(r.Port),
			Status: false,
			ID:     ID,
		})
	}()
	// new_chan.Range()
	// new_chan, err := t.Tracker.NewEndpoint(multicast.ReplayAll)
	return nil
}

// func (t *TrackerServer) Out(context.Context, r *tracking_protoc.StopRequest) (*tracking_protoc.Notify, error) {

// 	// return nil, nil
// }
