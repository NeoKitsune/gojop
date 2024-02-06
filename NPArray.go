package gojop

import (
	"encoding/binary"
	"strings"
)

var magic []byte = []byte{0x49, 0x68, 0x67, 0x1c}

func FromMsg() {
}

func PackMsg(name string, input []float32) []byte {
	var out []byte
	out = append(out, magic...)

	//4byte len
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint16(bs, uint16(len(input)))
	out = append(out, bs...)

	// Len of 2nd dim of array? arr.shape[0]
	binary.LittleEndian.PutUint16(bs, uint16(0))
	out = append(out, bs...)

	// Len of 1st dim of array? arr.shape[1]
	if i := uint16(len(input)); i <= 0 {
		binary.LittleEndian.PutUint16(bs, uint16(1))
	} else {
		binary.LittleEndian.PutUint16(bs, i)
	}
	out = append(out, bs...)

	// Len of 3rd dim of array? arr.shape[2]
	binary.LittleEndian.PutUint16(bs, uint16(0))
	out = append(out, bs...)

	// if type is not uint8
	out = append(out, 0x01)

	paddedName := name + strings.Repeat(" ", 128-len(name))
	out = append(out, []byte(paddedName)...)
	return out
}
