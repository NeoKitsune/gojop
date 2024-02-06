package gojop

import (
	"fmt"
	"net"
)

var (
	data []byte
	err  error
)

type SimEnv struct {
	Conn net.Conn
}

func (s *SimEnv) Connect(host string, port uint16) bool {
	if s.Conn != nil {
		return true
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	s.Conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		s.Conn = nil
		return false
	}
	go func() {
		data = make([]byte, 2048)
		for {
			if s.Conn == nil {
				break
			}
			n, err := s.Conn.Read(data)
			if err != nil {
				fmt.Println(err)
				break
			}
			if n > 0 {
				fmt.Printf("Server: %s\n", string(data[:n]))
			}
		}
	}()

	return true
}

/*
49 68 67 1c 00 00 00 00  00 00 00 00 01 00 00 00   Ihg..... ........
01 00 00 00 01 53 69 6d  45 6e 76 4d 61 6e 61 67   .....Sim EnvManag
65 72 2e 43 75 72 72 65  6e 74 2e 52 65 73 65 74   er.Curre nt.Reset
53 69 6d 45 6e 76 4e 6f  53 74 6f 70 20 20 20 20   SimEnvNo Stop
20 20 20 20 20 20 20 20  20 20 20 20 20 20 20 20
20 20 20 20 20 20 20 20  20 20 20 20 20 20 20 20
20 20 20 20 20 20 20 20  20 20 20 20 20 20 20 20
20 20 20 20 20 20 20 20  20 20 20 20 20 20 20 20
20 20 20 20 20 20 20 20  20 20 20 20 20 20 20 20
20 20 20 20 20
*/
func (s *SimEnv) Send(data []byte) {
	if s.Conn != nil {
		s.Conn.Write(data)
	}
}
