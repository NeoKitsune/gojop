package main

import (
	"fmt"

	. "github.com/NeoKitsune/gojop"
)

func main() {
	fmt.Println("JOP Test")

	speed := []float32{5}

	msg := PackMsg("ConveyorBelt._entityConveyorBelt0.setTargetSpeed", speed)
	PrintHex(msg)

	var s SimEnv
	s.Connect("127.0.0.1", 18189)
	s.Send(PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{}))
	s.Send(msg)

	fmt.Scanln()
	if s.Conn != nil {
		s.Conn.Close()
	}
}
