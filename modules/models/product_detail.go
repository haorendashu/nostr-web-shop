package models

import "xorm.io/xorm"

type ProductDetail struct {
	Id     string `xorm:"pk varchar(32)"`
	Pid    string `xorm:"notnull varchar(32) index(idx_product_detail)"`
	Status int    `xorm:"notnull"`
	Code   string `xorm:"notnull"`
	Name   string `xorm:"notnull"`
	Price  int    `xorm:"notnull"` // milisats, sats num * 1000
	Stock  int    `xorm:"notnull"`
}

func ProductDetailUpdateStock(id string, stock int, sessions ...*xorm.Session) error {
	s := getSession(sessions)
	sql := "update product_detail set stock = ? where id = ?"
	_, err := s.Exec(sql, stock, id)
	return err
}

func ProductDetailGet(id string, sessions ...*xorm.Session) *ProductDetail {
	o := &ProductDetail{}
	if objGet(sessions, o, "id = ?", id) {
		return o
	}
	return nil
}

func ProductDetailDel(pid string, sessions ...*xorm.Session) error {
	s := getSession(sessions)

	sql := "delete from product_detail where pid = ?"
	_, err := s.Exec(sql, pid)
	return err
}

func ProductDetailList(pid string, sessions ...*xorm.Session) []*ProductDetail {
	sql := "select * from product_detail where pid = ?"

	l := make([]*ProductDetail, 0)
	listQuery(sessions, &l, sql, pid)

	return l
}

func ProductDetailListByPids(pids []string, sessions ...*xorm.Session) []*ProductDetail {
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
