package common

import (
	"fmt"

	"github.com/gnufree/db_to_es/constant"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

func GetDBSession(serviceName string) (*mgo.Session, error) {
	c, err := constant.GetDBConfig(serviceName)
	if err != nil {
		return nil, err
	}

	dialInfo := &mgo.DialInfo{
		Addrs:    c.DBHosts,
		Database: c.DBName,
		Source:   c.Source,
		Username: c.DBUsername,
		Password: c.DBPassword,
	}

	return mgo.DialWithInfo(dialInfo)
}

func GetDB(serviceName string) (*gorm.DB, error) {
	c, err := constant.GetDBConfig(serviceName)
	if err != nil {
		return nil, err
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?authSource=admin&authMode=scram-sha1&charset=utf8&parseTime=True&loc=Local", c.DBUsername, c.DBPassword, c.DBHosts[0], c.DBName)
	return gorm.Open("mysql", dataSource)
}
