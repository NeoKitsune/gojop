package main

import (
	"fmt"
	"time"

	. "github.com/NeoKitsune/gojop"
	network "github.com/NeoKitsune/gojop/Network"
	"github.com/NeoKitsune/gojop/SimEnv"
	"github.com/charmbracelet/log"
)

func main() {
	log.Print("JOP Test")
	// log.SetLevel(log.DebugLevel)

	SimEnv.Connect("127.0.0.1", 18189)
	//Send(PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{}))
	simManager := SimEnvManager{}
	simManager.Reset()

	time.Sleep(5 * time.Second)
	log.Print(FindAll("ConveyorBelt"))
	speed := []float32{5}
	SimEnv.Send(network.PackMsg("ConveyorBelt._entityConveyorBelt0.setTargetSpeed", speed))

	fmt.Scanln()
	SimEnv.Disconnect()
}
