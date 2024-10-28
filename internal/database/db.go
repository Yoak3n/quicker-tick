package database

import (
	"log"
	"quicker-tick/internal/model"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	db = initSqlite("tick")
	if db == nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.TasksTable{}, &model.ActionsTable{}, &model.TagsTable{})
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("database connected")
}

func GetDB() *gorm.DB {
	return db
}
