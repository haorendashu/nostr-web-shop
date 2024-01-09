package models

import (
	"nostr-web-shop/modules/consts"
	"xorm.io/xorm"
)

type OrderProduct struct {
	Id            string `xorm:"pk varchar(32)"`
	OrderId       string `xorm:"notnull varchar(32) index(idx_order_product_oid)"`
	Pid           string `xorm:"notnull varchar(32) index(idx_order_product_pid)"` // this pid is product id.
	DetailId      string `xorm:"notnull varchar(32)"`
	Seller        string `xorm:"notnull varchar(64)"`
	Code          string `xorm:"notnull"`
	Name          string `xorm:"notnull"`
	Price         int    `xorm:"notnull"` // milisats, sats num * 1000
	Num           int    `xorm:"notnull"` // milisats, sats num * 1000
	Img           string // image
	PushCompleted int
}

func OrderProductListByOids(oids []string, sessions ...*xorm.Session) []*OrderProduct {
	args := make([]interface{}, 0)
	sql := "select * from order_product where order_id in ("
	for _, oid := range oids {
		args = append(args, oid)
		sql += "?,"
	}
	sql = sql[:len(sql)-1]
	sql += ")"

	l := make([]*OrderProduct, 0)
	listQuery(sessions, &l, sql, args...)

	return l
}

func OrderProductsNeededPush(sessions ...*xorm.Session) []*OrderProduct {
	l := make([]*OrderProduct, 0)
	sql := "select op.* from order_product op inner join `order` o on o.id = op.order_id where op.push_completed is null and o.pay_status = ? and o.status = ?"

	sqlArgs := make([]interface{}, 0)
	sqlArgs = append(sqlArgs, consts.PAY_STATUS_UNPAY)
	sqlArgs = append(sqlArgs, consts.DATA_STATUS_OK)

	listQuery(sessions, &l, sql, sqlArgs...)

	return l
}
