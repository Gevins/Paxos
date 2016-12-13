package src

type Instance struct {
	instance_id int
	value       string
	Proposer
	Acceptor
}

type Machine struct {
	id  int
	log string // local file path
}

func (m *Machine) init() {
}
