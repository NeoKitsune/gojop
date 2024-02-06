package main

import (
	"fmt"

	. "github.com/NeoKitsune/gojop"
)

func main() {
	fmt.Println("JOP Test")
	var s SimEnv
	s.Connect("127.0.0.1", 18189)
	s.Send("SimEnvManager.Current.ResetSimEnv")
	if s.Conn != nil {
		s.Conn.Close()
	}
}
