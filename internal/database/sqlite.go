package database

import (
	"fmt"
	"log"
	"quicker-tick/package/util"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initSqlite(name string) *gorm.DB {
	util.CreateDirNotExists("data/db")
	dsn := fmt.Sprintf("data/db/%s.db", name)
	sdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("database connected err:%v\n", err)
	}
	if err != nil {
		log.Printf("database connected err:%v\n", err)
	}
	return sdb
}
