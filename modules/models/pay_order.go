package models

import (
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/utils"
	"xorm.io/xorm"
)

type PayOrder struct {
	Id          string `xorm:"pk varchar(32)"`
	Oid         string `xorm:"notnull varchar(32) index(idx_pay_order)"`
	Price       int    `xorm:"notnull"` // milisats, sats num * 1000
	Pr          string `xorm:"notnull varchar(2048)"`
	VerifyUrl   string `xorm:"notnull varchar(512)"`
	CreatedAt   int64  `xorm:"notnull"`
	Status      int    `xorm:"notnull"`
	PayStatus   int    `xorm:"notnull"`
	ExpireTime  int64  `xorm:"notnull"`
	CheckedTime int64  `xorm:"notnull"` // The next check time
}

func PayOrderGet(id string, sessions ...*xorm.Session) *PayOrder {
	l := make([]*PayOrder, 0)
	sql := "select * from pay_order p where p.oid = ? and p.status = ? and p.expire_time > ? and p.pay_status in (?, ?)"

	sqlArgs := make([]interface{}, 0)
	sqlArgs = append(sqlArgs, id)
	sqlArgs = append(sqlArgs, consts.DATA_STATUS_OK)
	sqlArgs = append(sqlArgs, utils.NowInt64())
	sqlArgs = append(sqlArgs, consts.PAY_STATUS_PAIED)
	sqlArgs = append(sqlArgs, consts.PAY_STATUS_UNPAY)

	listQuery(sessions, &l, sql, sqlArgs...)

	if len(l) > 0 {
		for _, po := range l {
			if po.PayStatus == consts.PAY_STATUS_PAIED {
				return po
			}
		}

		return l[0]
	}

	return nil
}
