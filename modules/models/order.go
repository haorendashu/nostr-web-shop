package models

import (
	"nostr-web-shop/modules/consts"
	"xorm.io/xorm"
)

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
	Seller      string `xorm:"notnull varchar(64)"`
}

func OrderGet(id string, sessions ...*xorm.Session) *Order {
	o := &Order{}
	if objGet(sessions, o, "id = ?", id) {
		return o
	}
	return nil
}

func OrderListByBuyer(pubkey string, sessions ...*xorm.Session) []*Order {
	sql := "select * from `order` o where o.pubkey = ? and o.status = ? order by o.created_at desc"
	sqlArgs := make([]interface{}, 0)
	sqlArgs = append(sqlArgs, pubkey)
	sqlArgs = append(sqlArgs, consts.DATA_STATUS_OK)

	l := make([]*Order, 0)
	listQuery(sessions, &l, sql, sqlArgs...)

	return l
}
