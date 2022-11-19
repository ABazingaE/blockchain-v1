package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func uintToByte(num uint64) []byte {
	var buffer bytes.Buffer
	error := binary.Write(&buffer, binary.LittleEndian, num)
	if error != nil {
		fmt.Println("binary error:", error)
		return nil
	}

	return buffer.Bytes()
}
