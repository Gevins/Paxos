package src

import (
	"fmt"
	"strings"
)

type Acceptor struct {
	nextBal       int
	preVote       int
	propseMessage chan Message
	acceptMessage chan Message
	// Role
}

const (
	SEPARATOR = "|&|#|$|"
)

func (a *Acceptor) Promise(wait_time int) {
	promise_ok := 0
	wait_time_boom := time.After(wait_time * time.Millisecond)
	for {
		select {
		case c <- a.acceptMessage:
			// Just process the success stutus, or add the reject status if you want
			if c.value > a.nextBal {
				reply := Message{sender: a.id, value: fmt.Sprintf("%d%s%d", a.nextBal, SEPARATOR, a.preVote), state: 1, tag: 2}
				a.nextBal = c.value
				reply.Send(c.id)
			} else {
				reply := Message{sender: a.id, value: fmt.Sprintf("%d", a.nextBal), state: 0, tag: 2}
				reply.Send(c.id)
			}
		case <-wait_time_boom:
			return '2', nil
		default:
			log.Println("Wait the PROMISE message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func (a *Acceptor) Accepted(wait_time int) {
	promise_ok := 0
	wait_time_boom := time.After(wait_time * time.Millisecond)
	for {
		select {
		case c <- a.acceptMessage:
			// Just process the success stutus, or add the reject status if you want
			values := strings.Split(c.value, SEPARATOR)
			if len(values) == 2 && values[0] == a.nextBal {
				reply := Message{sender: a.id, value: fmt.Sprintf("%d%s%d", values[0], SEPARATOR, value[1]), state: 1, tag: 2}
				a.preVote = values[0]
				reply.Send(c.id)
			} else {
				reply := Message{sender: a.id, value: "", state: 0, tag: 2}
				reply.Send(c.id)
			}
		case <-wait_time_boom:
			return
		default:
			log.Println("Wait the ACCEPT message from Acceptors ... ")
			time.Sleep(100 * time.Microsecond)
		}
	}
}
