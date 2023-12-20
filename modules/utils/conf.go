package utils

import (
	"log"
)

var (
	CONFIG  = &Conf{}
	tokenIv []byte
)

func ParseConf() error {
	err := LoadJsonConf("config.json", CONFIG)
	if err != nil {
		log.Fatalf("config.json parse error %v", err)
	}

	return err
}

type Conf struct {
	Server string

	DBPath    string
	DBShowLog bool

	RootFiles   []string
	ShopPubkeys []string

	LoginEventU string
}
