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

func (a *Acceptor) Promise(m *Message, reply *Message) {
	reply.sender = a.id
	b := ValueToInt(m.value)
	if b > a.nextBal {
		reply.value = fmt.Sprintf("%d%s%d", a.nextBal, SEPARATOR, a.preVote)
		a.nextBal = b
		reply.state = 1
	} else {
		reply.value = fmt.Sprintf("%d", a.nextBal)
		reply.state = 0
	}
}

func (a *Acceptor) Accepted(m *Message, reply *Message) {
	reply.sender = a.id
	values := strings.Split(m.value, SEPARATOR)
	if len(values) == 2 && values[0] == a.nextBal {
		reply.value = fmt.Sprintf("VOTED:%d:%d", values[0], values[1])
		reply.state = 1
	} else {
		reply.state = 0
	}
}
