package entitybase

import (
	"slices"
	"strings"
)

var EntityMap = map[string][]string{}

// var EntityMap = map[string]classData{}

// type classData map[string]entityData
// type entityData map[string][]byte

func SyncIncomingData(names []string) {
	for _, name := range names {
		parts := strings.Split(name, ".")
		if len(parts) < 3 {
			continue
		}
		if parts[1][:1] == "_" {
			if val, ok := EntityMap[parts[0]]; ok {
				if slices.Contains(val, parts[1]) {
					continue
				}
			}
			EntityMap[parts[0]] = append(EntityMap[parts[0]], parts[1])
		}
	}
}
