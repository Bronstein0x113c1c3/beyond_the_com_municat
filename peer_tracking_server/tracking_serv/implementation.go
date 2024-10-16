package trackingserv

import (
	"tracking_server/tracking_protoc"
	"tracking_server/typing"
)

func first(r *tracking_protoc.IncomingRequest, s tracking_protoc.Tracking_AddServer, t *TrackerServer) string {

	ID := s.Context().Value(typing.T("channel")).(string)
	s.Send(&tracking_protoc.Notify{
		Status: true,
		Request: &tracking_protoc.Notify_Self{
			Self: &tracking_protoc.SelfNotify{
				ReceipentList: t.list(),
				Id:            ID,
			},
		},
	})
	t.add_someone(ID, typing.Notification{
		ID:     ID,
		Status: true,
		Addr:   r.Ip,
		Port:   int(r.Port)})
	t.Notify(&typing.Notification{
		Addr:   r.Ip,
		Port:   int(r.Port),
		Status: true,
		ID:     ID,
	})
	return ID
}
