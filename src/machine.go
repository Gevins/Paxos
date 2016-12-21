package src

import (
	//"net/rpc"
	"os"
	"net"
	"fmt"
	"log"
	"time"
)

// id, log should config in conf
type Machine struct {
	name    string
	log     string
	port    int
	msgs    chan Message
	istc    map[int]Instance
	istc_id int
	quorum  []string
	// Instance
}

func (m *Machine) init() {
	// Log init: create the log file if not exist
	if _, err := os.Stat(m.log); os.IsNotExist(err) {
		os.Create(m.log)
	}
	//machine := Machine{name: "", log: m.log, msgs: make(chan Message, 16), istc: make(map[int]Instance)}

	// Start the UDP service
	udpAddr, err := net.ResolveUDPAddr("udp4", m.name)
	CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	CheckError(err)
	defer conn.Close()
	for {
		var buf [512]byte
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err == nil {
			m.msgs <- ParseMessage(buf[:n])
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
	c := new(Message)
	for {
		select {
		case c <- m.msgs:
			// 0: client, 1: propose, 2:promise, 3:accept, 4:accepted
			tmp_istc := m.istc[c.instance_id]
			switch c.tag {
			case 0:
				// create new instance
				m.istc_id = m.istc_id + 1
				p := Proposer{lastTried: 0, clientMessage: c, promiseMessage: make(chan Message), acceptedMessage: make(chan Message)}
				a := Acceptor{nextBal: 0, preVote: 0}
				i := Instance{instance_id: m.istc_id, value: "", Proposer: p, Acceptor: a}
				m.istc[m.istc_id] = i
			case 1:
				//p := Proposer{lastTried: 0, clientMessage: c, promiseMessage: make(chan Message), acceptedMessage: make(chan Message)}
				//a := Acceptor{nextBal: 0, preVote: 0}
				//i := Instance{instance_id: m.istc_id, value: "", Proposer: p, Acceptor: a}
				//m.istc[c.instance_id].Acceptor.proposeMessage <- c
				tmp_istc.Acceptor.proposeMessage <- c
			case 2:
				tmp_istc.Proposer.promiseMessage <- c
			case 3:
				tmp_istc.Acceptor.acceptMessage <- c
			case 4:
				tmp_istc.Proposer.acceptedMessage <- c
			}
		default:
			log.Println("Wait the PROMISE message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
}
