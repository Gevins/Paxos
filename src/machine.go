package src

import (
	"net/rpc"
	"os"
)

// id, log should config in conf
type Machine struct {
	id   int
	name string
	port string
	log  string
	port int
	msgs chan Message
	istc map[int]Instance
	istc_id int
	// Instance
}

// type Instance struct {
// 	instance_id int
// 	value       string
// 	Proposer
// 	Acceptor
// }

func (m *Machine) init() {
	// Log init
	if _, err := os.Stat(m.log); os.IsNotExist(err) {
		os.Create(m.log)
	}
	// RPC register
	// id = , log = , instance_id = ,
	// Proposer{lastTried: int, clientMessage:, promise: make(chan Message), accepted: make(chan Message),Role:}
	// Acceptor{nextBal: int, preVote: int, Role{}}
	// Role{id: int, name: string, port: int, state: int}
	// pps := Proposer{}
	machine := Machine{id: 1, log: "path/to/file", msgs: make(chan Message, 16), istc: make(map[int]Instance))}

	// Start the UDP service
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	CheckError(err)
	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err == nil {
			count := 0
			for i := 0; i < len(buf); i++ {
				if buf[i] == '\n' {
					break
				}
				count++
			}
			m.msgs <- ParseMessage(buf[:count])
		}
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}


func (m *Machine) Run() {
	for {
		count++
		select {
		case c <-m.msgs:
			// 0: client, 1: propose, 2:promise, 3:accept, 4:accepted
			switch c.tag {
			case 0:
				// create new instance
				m.istc_id = m.istc_id + 1
				p := Proposer{lastTried: 0, clientMessage: c, promise: make(chan Message), accepted: make(chan Message)}
				a := Acceptor{nextBal: 0, preVote: 0}
				i := Instance{instance_id: m.istc_id, value: "", Proposer: p, Acceptor: a}
				m.istc[m.istc_id] = i
			case 1:
				m.istc[c.instance_id].Acceptor.propseMessage <- c
			case 2:
				m.istc[c.instance_id].Proposer.promiseMessage <- c
			case 3:
				m.istc[c.instance_id].Acceptor.acceptMessage <- c
			case 4:
				m.istc[c.instance_id].Proposer.acceptedMessage <- c
			}
			// Just process the success stutus, or add the reject status if you want
			if i.state == 1 {
				valus[promise_ok] = i.value
				promise_ok++
				if promise_ok > length {
					return true, values
				}
			}
		default:
			log.Println("Wait the PROMISE message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func processMessage() {
	
}