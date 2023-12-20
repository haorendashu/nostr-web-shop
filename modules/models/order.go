package models

type Order struct {
	Id          int64  `xorm:"pk autoincr"`
	Pubkey      string `xorm:"notnull varchar(64)"`
	UpdatedAt   int64  `xorm:"notnull"`
	CreatedAt   int64  `xorm:"notnull"`
	Status      int    `xorm:"notnull"`
	OrderStatus int    `xorm:"notnull"`
	PayStatus   int    `xorm:"notnull"`
	PaiedTime   int
	Price       int    `xorm:"notnull"` // milisats, sats num * 1000
	Comment     string `xorm:"varchar(512)"`
}
