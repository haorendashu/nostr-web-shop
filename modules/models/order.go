package models

type Order struct {
	Id          string `xorm:"pk"`
	Pubkey      string `xorm:"notnull varchar(64) index(idx_order)"`
	UpdatedAt   int64  `xorm:"notnull"`
	CreatedAt   int64  `xorm:"notnull"`
	Status      int    `xorm:"notnull"`
	OrderStatus int    `xorm:"notnull"`
	PayStatus   int    `xorm:"notnull"`
	PaiedTime   int
	Price       int    `xorm:"notnull"` // milisats, sats num * 1000
	Comment     string `xorm:"varchar(512)"`
}
