package models

import "xorm.io/xorm"

type ProductDetail struct {
	Id     int64  `xorm:"pk autoincr"`
	Pid    int64  `xorm:"notnull"`
	Status int    `xorm:"notnull"`
	Code   string `xorm:"notnull"`
	Name   string `xorm:"notnull"`
	Price  int    `xorm:"notnull"` // milisats, sats num * 1000
	Stock  int    `xorm:"notnull"`
}

func ProductDetailDel(pid int64, sessions ...*xorm.Session) error {
	s := getSession(sessions)

	sql := "delete from product_detail where pid = ?"
	_, err := s.Exec(sql, pid)
	return err
}

func ProductDetailList(pid int64, sessions ...*xorm.Session) []*ProductDetail {
	sql := "select * from product_detail where pid = ?"

	l := make([]*ProductDetail, 0)
	listQuery(sessions, &l, sql, pid)

	return l
}

func ProductDetailListByPids(pids []int64, sessions ...*xorm.Session) []*ProductDetail {
	args := make([]interface{}, 0)
	sql := "select * from product_detail where pid in ("
	for _, pid := range pids {
		args = append(args, pid)
		sql += "?,"
	}
	sql = sql[:len(sql)-1]
	sql += ")"

	l := make([]*ProductDetail, 0)
	listQuery(sessions, &l, sql, args...)

	return l
}
