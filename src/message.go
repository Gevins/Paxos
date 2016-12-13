package src

import (
	"flag"
	"strconv"
)

type Message struct {
	sender int
	value  string
	state  int
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
