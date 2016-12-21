package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

func main() {
	ok, str := test(10)
	fmt.Println(ok, str)

	cfg, err := ini.Load("src/config/config.ini")
	if err != nil {
		fmt.Printf(err.Error())
	}
	section, err := cfg.GetSection("dev")
	if err != nil {
		fmt.Printf("Can't Get the Section")
	}
	key, err := section.GetKey("quorum")

	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(key.String())
}

func test(len int) (bool, []int) {
	str := make([]int, len)
	str[0] = 1
	return true, str
}
