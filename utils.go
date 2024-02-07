package gojop

import "fmt"

func PrintHex(data []byte) {
	for i, p := range data {
		if i%16 == 0 && i != 0 {
			fmt.Printf(" | %s\n", stringHex(data[i-16:i]))
		} else if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02x ", p)
	}
	fmt.Println()
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
