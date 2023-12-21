package models

type OrderProduct struct {
	Id       string `xorm:"pk varchar(32)"`
	OrderId  string `xorm:"notnull varchar(32) index(idx_order_product_oid)"`
	Pid      string `xorm:"notnull varchar(32) index(idx_order_product_pid)"` // this pid is product id.
	DetailId string `xorm:"notnull varchar(32)"`
	Code     string `xorm:"notnull"`
	Name     string `xorm:"notnull"`
	Price    int    `xorm:"notnull"` // milisats, sats num * 1000
	Num      int    `xorm:"notnull"` // milisats, sats num * 1000
	Img      string // image
}
