package src

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Message struct {
	instance_id int
	sender      int
	value       string
	state       int
	tag         int // 0: client, 1: propose, 2:promise, 3:accept, 4:accepted
}

func ValueToInt(v string) int {
	flag.Parse()
	s := flag.Arg(v)
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}

func (m *Message) String() string {
	return fmt.Sprintf("%d %s %d", m.sender, m.value, m.state)
}

func ParseMessage(m []byte) Message {
	strs := strings.Split(string(m), "\t")
	return Message{sender: strs[0], value: strs[1], state: strs[2]}
}
