package src

import (
	"fmt"
	"log"
	//"net/rpc"
	"time"
)

const (
	SEPARATOR = "|&|#|$|"
)

type Proposer struct {
	lastTried       int
	clientMessage   string
	promiseMessage  chan Message
	acceptedMessage chan Message
	Role
}

// Prepare call Prepare of Acceptor, return the prepare requese is or not success and the promise message.
// Quorum should be define in config
// Wait time should be init for get the promise
func (p *Proposer) Propose(wait_time int) (byte, map[int]string) {
	p.lastTried = p.lastTried + 1
	send := Message{sender: p.id, value: p.lastTried, state: 1, tag: 1}

	length := len(p.Role.quorum)
	for i := 0; i < length; i++ {
		log.Println("Send PREPARE request to:", p.Role.quorum[i])
		// go rpcCall(p.Role.quorum[i], "Acceptor.Promise:", send, p.promise)
		send.Send(p.Role.quorum[i])
	}

	promise_ok := 0
	wait_time_boom := time.After(wait_time * time.Millisecond)
	values := make(map[int]string)
	count := 0
	c := new(Message)
	for {
		select {
		case c <- p.promiseMessage:
			// Just process the success stutus, or add the reject status if you want
			if c.state == 1 {
				values[c.sender] = c.value
				promise_ok++
				if promise_ok > length {
					// start Accept step
					return '0', values
				}
			} else {
				count++
				if count == length {
					// Rejected
					return '1', nil
				}
			}
		case <-wait_time_boom:
			// time out, should be propose again
			// time.Sleep(time.Second)
			// go p.Propose(wait_time)
			return '2', nil
		default:
			log.Println("Wait the PROMISE message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func (p *Proposer) Accept(wait_time int, value string) {
	send := Message{sender: p.id, value: fmt.Sprintf("%d%s%d", p.lastTried, value), state: 1, tag: 3}
	length := len(p.Role.quorum)
	for i := 0; i < length; i++ {
		// go rpcCall(quorum[i], "Acceptor.Accepted:", send, p.accepted)
		send.Send(p.Role.quorum[i])
	}

	accepted_ok := 0
	wait_time_boom := time.After(wait_time * time.Millisecond)
	count := 0
	values := make(map[int]string)
	c := new(Message)
	for {
		select {
		case c <- p.promiseMessage:
			if c.state == 1 {
				values[accepted_ok] = c.value
				accepted_ok++
				if accepted_ok > length {
					// aceess next step
					// return true, values
					fmt.Println("scceess")
					return
				}
			} else {
				count++
				if count == length {
					// Rejected
					return
				}
			}
		case <-wait_time_boom:
			return
		default:
			log.Println("Wait the ACCEPTED message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
	return
}

// func rpcCall(server string, process string, send *Message, c chan Message) {
// 	client, err := rpc.Dial("tcp", server)
// 	if err != nil {
// 		log.Fatal(process, err)
// 	}
// 	var reply Message
// 	client.Call(process, send, &reply)
// 	c <- reply
// }

func (p *Proposer) run() {
	wait_time := time.Millisecond * 10
	// process time out
	for {
		go func() {
			state, values := p.Propose(wait_time)
			if state == '0' {
				tmp_key, tmp_value := 0, ""
				for key, value := range values {
					if key > tmp_key {
						tmp_value = value
					}
				}
				p.Accept(wait_time, tmp_value)

			} else if state == '1' {
				return
			} else {
				// time out
				return
			}
		}()
	}
	// for {
	// 	p.Accept(wait_time, value)
	// }
}
