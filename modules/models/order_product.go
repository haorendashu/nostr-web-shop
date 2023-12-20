package models

type OrderProduct struct {
	Id    int64  `xorm:"pk autoincr"`
	Pid   int64  `xorm:"notnull"` // this pid is product id.
	Code  string `xorm:"notnull"`
	Name  string `xorm:"notnull"`
	Price int    `xorm:"notnull"` // milisats, sats num * 1000
	Num   int    `xorm:"notnull"` // milisats, sats num * 1000
	Img   string // image
}
