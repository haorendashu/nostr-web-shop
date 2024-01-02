package models

import "xorm.io/xorm"

type OrderProduct struct {
	Id       string `xorm:"pk varchar(32)"`
	OrderId  string `xorm:"notnull varchar(32) index(idx_order_product_oid)"`
	Pid      string `xorm:"notnull varchar(32) index(idx_order_product_pid)"` // this pid is product id.
	DetailId string `xorm:"notnull varchar(32)"`
	Seller   string `xorm:"notnull varchar(64)"`
	Code     string `xorm:"notnull"`
	Name     string `xorm:"notnull"`
	Price    int    `xorm:"notnull"` // milisats, sats num * 1000
	Num      int    `xorm:"notnull"` // milisats, sats num * 1000
	Img      string // image
}

func OrderProductListByOids(pids []string, sessions ...*xorm.Session) []*OrderProduct {
	args := make([]interface{}, 0)
	sql := "select * from order_product where order_id in ("
	for _, pid := range pids {
		args = append(args, pid)
		sql += "?,"
	}
	sql = sql[:len(sql)-1]
	sql += ")"

	l := make([]*OrderProduct, 0)
	listQuery(sessions, &l, sql, args...)

	return l
}
