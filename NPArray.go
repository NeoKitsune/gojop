package gojop

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/log"
)

var (
	magic      []byte = []byte{0x49, 0x68, 0x67, 0x1c}
	nameLen           = 128
	preHeader         = 17
	headerSize        = nameLen + preHeader
)

func FromMsg(msg []byte) {
	packets := bytes.Split(msg, magic)

	for _, packet := range packets {
		if len(packet) < 4 {
			continue
		}
		msgLen := binary.LittleEndian.Uint16(packet[:4])
		w := binary.LittleEndian.Uint16(packet[4:8])
		h := binary.LittleEndian.Uint16(packet[8:12])
		c := binary.LittleEndian.Uint16(packet[12:16])
		dt := int8(packet[16])

		if len(packet) < headerSize {
			PrintHex(packet)
			fmt.Println("ERROR")
			log.Error("Bad Packet From Server")
			continue
		}
		name := strings.TrimSpace(string(packet[preHeader:headerSize]))

		PrintHex(packet)
		log.Debugf("Name: %s\n\tLen: %d w: %d h: %d c: %d dt: %d\n", name, msgLen, w, h, c, dt)
	}
}

func PackMsg(name string, input []float32) []byte {
	log.Infof("CMD: %s Values: %v", name, input)
	var out []byte
	out = append(out, magic...)

	inputBytes := Float32Arrybits(input)

	//4byte len
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint16(bs, uint16(len(inputBytes)))
	out = append(out, bs...)

	// Len of 2nd dim of array? arr.shape[0]
	binary.LittleEndian.PutUint16(bs, uint16(len(input)))
	out = append(out, bs...)

	// Len of 1st dim of array? arr.shape[1]
	if i := uint16(len(input)); i <= 0 {
		binary.LittleEndian.PutUint16(bs, uint16(1))
	} else {
		binary.LittleEndian.PutUint16(bs, i)
	}
	out = append(out, bs...)

	// Len of 3rd dim of array? arr.shape[2]
	binary.LittleEndian.PutUint16(bs, uint16(1))
	out = append(out, bs...)

	// if type is not uint8
	out = append(out, 0x01)

	paddedName := name + strings.Repeat(" ", nameLen-len(name))
	out = append(out, []byte(paddedName)...)

	out = append(out, inputBytes...)
	return out
}

func Float32Arrybits(slice []float32) []byte {
	var out []byte
	for _, i := range slice {
		out = binary.LittleEndian.AppendUint32(out, math.Float32bits(i))
	}
	return out
}
