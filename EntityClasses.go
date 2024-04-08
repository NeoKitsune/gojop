package gojop

import (
	entitybase "github.com/NeoKitsune/gojop/EntityBase"
	network "github.com/NeoKitsune/gojop/Network"
	"github.com/NeoKitsune/gojop/SimEnv"
)

func FindAll(entityType string) []string {
	if list, ok := entitybase.EntityMap[entityType]; ok {
		return list
	}
	return []string{}
}

type SimEnvManager struct{}

func (sim *SimEnvManager) Reset() {
	SimEnv.Send(network.PackMsg("SimEnvManager.Current.ResetSimEnvNoStop", []float32{}))
}
