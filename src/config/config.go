package config

import (
	"github.com/go-ini/ini"
	"fmt"
)

func getValue(key string) string {
	cfg, err := ini.Load("src/config/config.ini")
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	section, err0 := cfg.GetSection("dev")
	if err0 != nil {
		fmt.Printf(err0.Error())
		return ""
	}
	key, err1 := section.GetKey("quorum")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return ""
	}
	return key.String()
}