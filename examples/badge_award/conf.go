package main

import (
	"log"
	"nostr-web-shop/modules/utils"
)

var (
	CONFIG  = &Conf{}
	tokenIv []byte
)

func ParseConf() error {
	err := utils.LoadJsonConf("config.json", CONFIG)
	if err != nil {
		log.Fatalf("config.json parse error %v", err)
	}

	return err
}

type Conf struct {
	Server    string
	RootFiles []string

	PushKey      string
	PrivateKey   string
	ProductCodes []string
	BadgeAIds    []string
	NoticeRelays []string
}
