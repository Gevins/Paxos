package src

import (
	"flag"
	"fmt"
	"net"
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
	return fmt.Sprintf("%d%s%d%s%d%s%s%s%d%s%d", m.instance_id, m.sender, SEPARATOR, m.value, SEPARATOR, m.state, SEPARATOR, m.tag)
}

func ParseMessage(m []byte) Message {
	strs := strings.Split(string(m), SEPARATOR)
	return Message{instance_id: strs[0], sender: strs[1], value: strs[2], state: strs[3], tag: strs[4]}
}

func (m *Message) Send(rcv string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", rcv)
	CheckError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte(m.String()))
	checkError(err)
}
