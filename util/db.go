package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

func GetConnection() (db *gorm.DB, err error) {

	conf, err := ini.Load("conf/config.ini")
	if err != nil {
		return nil, err
	}

	dbType := conf.Section("db").Key("type").String()
	DBMS := conf.Section(dbType).Key("dbms").String()
	USER := conf.Section(dbType).Key("user").String()
	PASS := conf.Section(dbType).Key("pass").String()
	PROTOCOL := conf.Section(dbType).Key("protocol").String()
	DBNAME := conf.Section(dbType).Key("db_name").String()

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
