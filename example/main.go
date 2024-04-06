package main

import (
	"fmt"

	. "github.com/NeoKitsune/gojop"
	"github.com/charmbracelet/log"
)

func main() {
	log.Print("JOP Test")
	//log.SetLevel(log.DebugLevel)

	var sim SimEnv
	sim.Connect("127.0.0.1", 18189)
	sim.Send(PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{}))

	speed := []float32{5}
	sim.Send(PackMsg("ConveyorBelt._entityConveyorBelt0.setTargetSpeed", speed))

	fmt.Scanln()
	sim.Disconnect()
}
