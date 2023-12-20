package models

type PayOrder struct {
	Id          string `xorm:"pk"`
	Oid         string `xorm:"notnull index(idx_pay_order)"`
	Price       int    `xorm:"notnull"` // milisats, sats num * 1000
	Pr          string `xorm:"notnull varchar(2048)"`
	VerifyUrl   string `xorm:"notnull varchar(512)"`
	CreatedAt   int64  `xorm:"notnull"`
	Status      int    `xorm:"notnull"`
	ExpireTime  int64  `xorm:"notnull"`
	CheckedTime int64  `xorm:"notnull"` // The next check time
}
