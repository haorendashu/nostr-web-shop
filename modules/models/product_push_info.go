package models

import (
	"nostr-web-shop/modules/consts"
	"xorm.io/xorm"
)

type ProductPushInfo struct {
	Id           string `xorm:"pk varchar(32)"`
	Pid          string `xorm:"notnull varchar(32) index(idx_product_push_info)"`
	Status       int    `xorm:"notnull"`
	NoticePubkey string `xorm:"notnull"`
	PushAddress  string `xorm:"notnull"`
	PushKey      string `xorm:"notnull"`
	PushType     int    `xorm:"notnull"` // 1 api push, 2 web push
}

func ProductPushInfoGetByCode(pubkey, code string, sessions ...*xorm.Session) *ProductPushInfo {
	sql := "select ppi.* from product p inner join product_detail pd on p.id = pd.pid inner join product_push_info ppi on ppi.pid = p.id where p.status = ? and ppi.status = ? and pd.status = ? and p.pubkey = ? and pd.code = ?"
	sqlArgs := []interface{}{consts.DATA_STATUS_OK, consts.DATA_STATUS_OK, consts.DATA_STATUS_OK, pubkey, code}

	l := make([]*ProductPushInfo, 0)
	listQuery(sessions, &l, sql, sqlArgs...)

	if len(l) > 0 {
		return l[0]
	}

	return nil
}

func ProductPushInfoGet(pid string, sessions ...*xorm.Session) *ProductPushInfo {
	o := &ProductPushInfo{}
	if objGet(sessions, o, "pid = ?", pid) {
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

func ProductPushInfoListByPids(pids []string, sessions ...*xorm.Session) []*ProductPushInfo {
	args := make([]interface{}, 0)
	sql := "select * from product_push_info where pid in ("
	for _, oid := range pids {
		args = append(args, oid)
		sql += "?,"
	}
	sql = sql[:len(sql)-1]
	sql += ")"

	l := make([]*ProductPushInfo, 0)
	listQuery(sessions, &l, sql, args...)

	return l
}
