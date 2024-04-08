package utils

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func PrintHex(data []byte) {
	log.Debug(PrettyHex(data))
}

func PrettyHex(data []byte) string {
	out := "\n"
	for i, p := range data {
		if i%16 == 0 && i != 0 {
			out += fmt.Sprintf(" | %s\n", stringHex(data[i-16:i]))
		} else if i%8 == 0 && i != 0 {
			out += " "
			//fmt.Print(" ")
		}
		out += fmt.Sprintf("%02x ", p)
	}
	//fmt.Println()
	out += "\n"
	return out
}

func stringHex(hs []byte) string {
	var out string
	for i := range hs {
		if hs[i] < 0x20 || hs[i] >= 0x7f {
			out += "."
		} else {
			out += string(hs[i])
		}
	}
	return out
}
