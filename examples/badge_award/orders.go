package main

import (
	"log"
	"nostr-web-shop/modules/utils"
	"os"
)

var ORDERS = make([]string, 0)

const ORDERS_FILE_NAME = "orders.txt"

func loadOrders() {
	os, err := utils.LoadPlainArray(ORDERS_FILE_NAME)
	if err == nil {
		ORDERS = os
	} else {
		log.Fatalf("%s read error %v", ORDERS_FILE_NAME, err)
	}
}

func addOrder(oid string) {
	file, err := os.OpenFile(ORDERS_FILE_NAME, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open orders.txt fail %v", err)
		return
	}
	defer file.Close()

	file.WriteString(oid)
	file.WriteString("\n")
	file.Sync()

	ORDERS = append(ORDERS, oid)
}

func checkOrder(oid string) bool {
	for _, o := range ORDERS {
		if o == oid {
			return true
		}
	}
	return false
}
