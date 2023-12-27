package dtos

import "nostr-web-shop/modules/models"

type OrderDto struct {
	Id          string
	Pubkey      string
	CreatedAt   int64
	Status      int
	OrderStatus int
	PayStatus   int
	PaiedTime   int
	Price       int
	Lnwallet    string
	Comment     string
	Seller      string

	Skus []*models.OrderProduct
}
