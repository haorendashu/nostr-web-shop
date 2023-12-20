package models

type OrderProduct struct {
	Id    string `xorm:"pk"`
	Pid   string `xorm:"notnull index(idx_order_product)"` // this pid is product id.
	Code  string `xorm:"notnull"`
	Name  string `xorm:"notnull"`
	Price int    `xorm:"notnull"` // milisats, sats num * 1000
	Num   int    `xorm:"notnull"` // milisats, sats num * 1000
	Img   string // image
}
