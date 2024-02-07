package main

import (
	"fmt"

	. "github.com/NeoKitsune/gojop"
)

func main() {
	fmt.Println("JOP Test")

	msg := PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{})
	PrintHex(msg)

	var s SimEnv
	s.Connect("127.0.0.1", 18189)
	s.Send(PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{}))

	fmt.Scanln()
	if s.Conn != nil {
		s.Conn.Close()
	}
}
