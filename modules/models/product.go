package models

import (
	"nostr-web-shop/modules/consts"
	"xorm.io/xorm"
)

type Product struct {
	Id        int64  `xorm:"pk autoincr"`
	Pubkey    string `xorm:"notnull varchar(64)"`
	UpdatedAt int64  `xorm:"notnull"`
	CreatedAt int64  `xorm:"notnull"`
	Status    int    `xorm:"notnull"`
	Name      string `xorm:"notnull"`
	Imgs      string `xorm:"notnull"` // images join ,
	Des       string `xorm:"notnull varchar(1024)"`
	Content   string `xorm:"notnull TEXT"` // html content
	Price     int    `xorm:"notnull"`      // milisats, sats num * 1000
	Lnwallet  string `xorm:"notnull varchar(128)"`
}

func ProductGet(id int64, sessions ...*xorm.Session) *Product {
	o := &Product{}
	if objGet(sessions, o, "id = ?", id) {
		return o
	}
	return nil
}

func ProductDel(pid int64, sessions ...*xorm.Session) error {
	s := getSession(sessions)

	sql := "delete from product where id = ?"
	_, err := s.Exec(sql, pid)
	return err
}

func ProductList(args *ProductQueryArgs, sessions ...*xorm.Session) []*Product {
	sql := "select * from product where status = ? "
	sqlArgs := make([]interface{}, 0)
	sqlArgs = append(sqlArgs, consts.DATA_STATUS_OK)

	if args.Pubkey != "" {
		sql += "and pubkey = ? "
		sqlArgs = append(sqlArgs, args.Pubkey)
	}
	sql += "order by updated_at desc"

	l := make([]*Product, 0)
	listQuery(sessions, &l, sql, sqlArgs...)

	return l
}

type ProductQueryArgs struct {
	Pubkey string
}
