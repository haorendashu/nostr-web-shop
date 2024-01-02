package models

import (
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/dtos"
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
	PaidTime    int64
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

func OrderListByBuyer(pubkey string, sessions ...*xorm.Session) []*Order {
	sql := "select * from `order` o where o.pubkey = ? and o.status = ? order by o.created_at desc"
	sqlArgs := make([]interface{}, 0)
	sqlArgs = append(sqlArgs, pubkey)
	sqlArgs = append(sqlArgs, consts.DATA_STATUS_OK)

	l := make([]*Order, 0)
	listQuery(sessions, &l, sql, sqlArgs...)

	return l
}

func OrderApiList(seller, code string, since int64, sessions ...*xorm.Session) []*dtos.ApiOrderDto {
	sql := "select o.seller, op.code, op.id as orderProductId, o.pubkey as buyer, op.num, o.comment, o.pay_status as payStatus, o.paid_time as paidTime from `order` o inner join order_product op on o.id = op.order_id where o.status = ? and o.seller = ? and o.pay_status = ? and op.code = ? and o.paid_time > ? order by o.paid_time asc"
	sqlArgs := []interface{}{consts.DATA_STATUS_OK, seller, consts.PAY_STATUS_PAIED, code, since}

	l := make([]*dtos.ApiOrderDto, 0)
	listQuery(sessions, &l, sql, sqlArgs...)

	return l
}

func OrderApiGet(orderProductId string, sessions ...*xorm.Session) *dtos.ApiOrderDto {
	sql := "select o.seller, op.code, op.id as orderProductId, o.pubkey as buyer, op.num, o.comment, o.pay_status as payStatus, o.paid_time as paidTime from `order` o inner join order_product op on o.id = op.order_id where o.status = ? and op.id = ?"
	sqlArgs := []interface{}{consts.DATA_STATUS_OK, orderProductId}

	l := make([]*dtos.ApiOrderDto, 0)
	listQuery(sessions, &l, sql, sqlArgs...)

	if len(l) > 0 {
		return l[0]
	}

	return nil
}
