package src

import (
	"fmt"
	// "log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Role struct {
	id int
	// name  string
	// port  int
	// state int // 0:stop, 1:running, 2:stoping
	quorum []string
}

func (r *Role) Prepare() {

}

type Manner interface {
	Start()
	Stop()
	Send(m *Message)
	Recieve()
}

func (r *Role) Start() {
	rpc.Register(*r)
	adrr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", r.port))
	checkError(err)
	listener, err := net.ListenTCP("tcp", adrr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			if r.state == 2 {
				stop(listener)
			} else {
				continue
			}
		}
		fmt.Println("Accept a news!")
		go rpc.ServeConn(conn)
	}
}

func stop(l *net.TCPListener) {
	l.Close()
	os.Exit(0)
}

func (r *Role) Stop() {
	r.state = 2
}

func (r *Role) handleclient(conn net.Conn) {
	defer conn.Close()
	// Todo: add the code for rpc process

}

func (r *Role) send(to string, m Message) {
	udpAddr, err := net.ResolveUDPAddr("udp4", to)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte(m.String()))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
