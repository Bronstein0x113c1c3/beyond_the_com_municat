package main

import (
	"encoding"
	"io"

	"github.com/google/uuid"
)

/*



 */

var (
	_ io.WriterTo                = (*Packet)(nil)
	_ io.ReaderFrom              = (*Packet)(nil)
	_ encoding.BinaryMarshaler   = (*Packet)(nil)
	_ encoding.BinaryUnmarshaler = (*Packet)(nil)
)

type PacketType byte

const (
	PacketJoin PacketType = iota + 1
	PacketAudio
	PacketLeave
)

type Packet struct {
	Type    PacketType
	ID      uuid.UUID
	Payload PacketPayload
}

var (
	_ io.WriterTo                = (*PacketPayload)(nil)
	_ io.ReaderFrom              = (*PacketPayload)(nil)
	_ encoding.BinaryMarshaler   = (*PacketPayload)(nil)
	_ encoding.BinaryUnmarshaler = (*PacketPayload)(nil)
)

type PacketPayload struct {
	data []byte
}

func (p *PacketPayload) Size() int16 {
	return int16(len(p.data) + 2)
}
