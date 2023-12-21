package models

import "xorm.io/xorm"

type ProductPushInfo struct {
	Id           string `xorm:"pk varchar(32)"`
	Pid          string `xorm:"notnull varchar(32) index(idx_product_push_info)"`
	Status       int    `xorm:"notnull"`
	NoticePubkey string `xorm:"notnull"`
	PushAddress  string `xorm:"notnull"`
	PushKey      string `xorm:"notnull"`
	PushType     int    `xorm:"notnull"` // 1 api push, 2 web push
}

func ProductPushInfoGet(id string, sessions ...*xorm.Session) *ProductPushInfo {
	o := &ProductPushInfo{}
	if objGet(sessions, o, "pid = ?", id) {
		return o
	}
	return nil
}

func ProductPushInfoDel(pid string, sessions ...*xorm.Session) error {
	s := getSession(sessions)

	sql := "delete from product_push_info where pid = ?"
	_, err := s.Exec(sql, pid)
	return err
}
