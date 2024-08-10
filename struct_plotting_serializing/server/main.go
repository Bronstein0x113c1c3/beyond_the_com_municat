package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

/*
encoding scheme:
	type + ID + msg_len + msg = 1 + 16 + 2 + variable_length
	max_size = 1600



*/

func main() {

	buf := new(bytes.Buffer)
	fmt.Println(buf.Available())
	// fmt.Println("asfdsdf")
	new_uuid := uuid.New()
	// uuid.ParseBytes()
	bytes, _ := new_uuid.MarshalBinary()
	fmt.Println(bytes)
	// bytes, _ := new_uuid.MarshalBinary()
	fmt.Printf("%v \n", len(new_uuid))
	binary.Write(buf, binary.LittleEndian, new_uuid)
	fmt.Printf("%v \n", buf.Bytes())
	// binary.Read()

}
