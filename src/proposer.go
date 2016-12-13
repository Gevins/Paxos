package src

import (
	"fmt"
	// log "github.com/cihub/seelog"
	"log"
	"net/rpc"
	"time"
)

const (
	SEPARATOR = "|&|#|$|"
)

type Proposer struct {
	lastTried     int
	clientMessage chan Message
	promise       chan Message
	accepted      chan Message
	Role
}

type Instance struct {
	instance_id int
	value       string
	Proposer
	Acceptor
}

// Prepare call Prepare of Acceptor, return the prepare requese is or not success and the promise message.
// Quorum should be define in config
// Wait time should be init for get the promise
func (p *Proposer) Prepare(quorum []string, length int, wait_time int) (bool, []string) {
	p.lastTried = lastTried + 1
	send := Message{sender: p.id, value: p.lastTried, state: 1}

	for i := 0; i < len(quorum); i++ {
		log.Println("Send PREPARE request to:", quorum[i])
		go rpcCall(quorum[i], "Acceptor.Prepare:", send, p.promise)
	}

	promise_ok := 0
	wait_time_boom := time.After(wait_time * time.Millisecond)
	values := make([]string, length)
	count := 0
	for {
		count++
		select {
		case <-p.promise:
			// Just process the success stutus, or add the reject status if you want
			if i.state == 1 {
				valus[promise_ok] = i.value
				promise_ok++
				if promise_ok > length {
					return true, values
				}
			}
		case <-wait_time_boom:
			return false, values
		default:
			log.Println("Wait the PROMISE message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
	if count == len(quorum) {
		return false, values
	}

}

func (p *Proposer) Accept(quorum []string, length int, wait_time int, value string) (bool, []string) {
	send := Message{sender: p.id, value: fmt.Sprintf("%d%s%d", p.lastTried, value), state: 1}
	for i := 0; i < len(quorum); i++ {
		go rpcCall(quorum[i], "Acceptor.Accept:", send, p.accepted)
	}

	accepted_ok := 0
	wait_time_boom := time.After(wait_time * time.Millisecond)
	values := make([]string, length)
	count := 0
	for i := range p.accepted {
		count++
		if i.state == 1 {
			valus[accepted_ok] = i.value
			accepted_ok++
			if accepted_ok > length {
				return true, values
			}
		}
		if count == len(quorum) || time.Now().Sub(wait_time_begin).Nanoseconds()/1000 > wait_time {
			break
		}
	}
	for {
		select {
		case <-p.promise:
			if i.state == 1 {
				valus[accepted_ok] = i.value
				accepted_ok++
				if accepted_ok > length {
					return true, values
				}
			}
		case <-wait_time_boom:
			return false, values
		default:
			log.Println("Wait the ACCEPTED message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
	return false, values
}

func rpcCall(server string, process string, send *Message, c chan Message) {
	client, err := rpc.Dial("tcp", server)
	if err != nil {
		log.Fatal(process, err)
	}
	var reply Message
	client.Call(process, send, &reply)
	c <- reply
}
