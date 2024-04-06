package gojop

import (
	network "github.com/NeoKitsune/gojop/Network"
	"github.com/NeoKitsune/gojop/SimEnv"
)

type SimEnvManager struct{}

func (sim *SimEnvManager) Reset() {
	SimEnv.Send(network.PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{}))
}
