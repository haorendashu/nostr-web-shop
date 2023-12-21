package models

import "xorm.io/xorm"

type Order struct {
	Id          string `xorm:"pk varchar(32)"`
	Pubkey      string `xorm:"notnull varchar(64) index(idx_order)"`
	UpdatedAt   int64  `xorm:"notnull"`
	CreatedAt   int64  `xorm:"notnull"`
	Status      int    `xorm:"notnull"`
	OrderStatus int    `xorm:"notnull"`
	PayStatus   int    `xorm:"notnull"`
	PaiedTime   int
	Price       int    `xorm:"notnull"` // milisats, sats num * 1000
	Lnwallet    string `xorm:"notnull"`
	Comment     string `xorm:"varchar(512)"`
}

func OrderGet(id string, sessions ...*xorm.Session) *Order {
	o := &Order{}
	if objGet(sessions, o, "id = ?", id) {
		return o
	}
	return nil
}
