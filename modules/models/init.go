package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"nostr-web-shop/modules/utils"
	"xorm.io/xorm"
)

var (
	engine *xorm.Engine
)

func Init() {
	var err error
	engine, err = xorm.NewEngine("mysql", utils.CONFIG.DBPath)
	if err != nil {
		log.Printf("db init error %v", err)
	}

	engine.ShowSQL(utils.CONFIG.DBShowLog)
}

func Sync() {
	err := engine.Sync(&Order{}, &OrderProduct{}, &PayOrder{}, &Product{}, &ProductDetail{}, &ProductPushInfo{})
	if err != nil {
		log.Fatalf("engine.Sync error %v", err)
	}
}

func Begin() *xorm.Session {
	session := engine.NewSession()
	session.Begin()
	return session
}

func Commit(session *xorm.Session) {
	session.Commit()
}

func Rollback(session *xorm.Session) {
	session.Rollback()
}

func ObjInsert(o interface{}, sessions ...*xorm.Session) bool {
	s := getSession(sessions)
	_, err := s.InsertOne(o)
	if err != nil {
		log.Printf("obj save error %v", err)
		return false
	}

	return true
}

func ObjUpdate(id interface{}, o interface{}, sessions ...*xorm.Session) bool {
	s := getSession(sessions)
	_, err := s.ID(id).Update(o)
	if err != nil {
		log.Printf("ObjUpdate error %v", err)
		return false
	}

	return true
}

func objGet(sessions []*xorm.Session, result interface{}, whereSql string, args ...interface{}) bool {
	s := getSession(sessions)
	has, err := s.Where(whereSql, args...).Get(result)
	if err != nil {
		log.Printf("ObjGet sql %s error %v", whereSql, err)
	}

	return has
}

func listQuery(sessions []*xorm.Session, result interface{}, sql string, args ...interface{}) interface{} {
	s := getSession(sessions)
	err := s.SQL(sql, args...).Find(result)
	if err != nil {
		log.Printf("listQuery sql %s error %v", sql, err)
	}
	return result
}

func listCount(sessions []*xorm.Session, sql string, args ...interface{}) int64 {
	s := getSession(sessions)
	count, err := s.SQL(sql, args...).Count()
	if err != nil {
		log.Printf("listCount sql %s error %v", sql, err)
	}
	return count
}

func getSession(sessions []*xorm.Session) *xorm.Session {
	if sessions != nil && len(sessions) > 0 {
		return sessions[0]
	}

	return engine.NewSession()
}
