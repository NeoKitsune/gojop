package main

import (
	"fmt"

	. "github.com/NeoKitsune/gojop"
)

func main() {
	fmt.Println("JOP Test")

	msg := PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{})
	for i, p := range msg {
		if i%16 == 0 {
			fmt.Println()
		} else if i%8 == 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02x ", p)
	}

	var s SimEnv
	s.Connect("127.0.0.1", 18189)
	s.Send(PackMsg("SimEnvManager.Current.ResetSimEnv", []float32{}))
	if s.Conn != nil {
		s.Conn.Close()
	}
}
