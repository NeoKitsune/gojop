package SimEnv

import (
	"fmt"
	"net"

	network "github.com/NeoKitsune/gojop/Network"
	"github.com/NeoKitsune/gojop/utils"
	"github.com/charmbracelet/log"
)

var (
	data []byte
	err  error
)

var conn net.Conn

func Connect(host string, port uint16) bool {
	if conn != nil {
		return true
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		//fmt.Println(err)
		log.Error(err)
		conn = nil
		return false
	}
	log.Info("Connected to Sim")
	go func() {
		data = make([]byte, 4096)
		for {
			if conn == nil {
				break
			}
			n, err := conn.Read(data)
			if err != nil {
				// fmt.Println(err)
				log.Error(err)
				break
			}
			if n > 0 {
				network.FromMsg(data[:n])
				//fmt.Printf("Server: %s\n", string(data[:n]))
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
func Send(data []byte) {
	if conn != nil {
		utils.PrintHex(data)
		conn.Write(data)
	}
}

func Disconnect() {
	if conn == nil {
		log.Warn("Trying to close nil connection")
		return
	}
	conn.Close()
	log.Info("Disconnect from Sim")

}
