package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
)

func init() {

	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/layuicmsCarRent?charset=utf8")
	orm.RegisterModel(new(User), new(Menu))
	orm.RunSyncdb("default", false, true)

	initSession()
}

var globalSessions *session.Manager

func initSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmpsession",
	}
	globalSessions, _ = session.NewManager("file", sessionConfig)
	go globalSessions.GC()
}
