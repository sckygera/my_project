package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jinzhu/gorm"
	"log"
	"my_project/model/Book"
	"my_project/model/author"
	"time"
)

var dbase *gorm.DB

type Client interface {
	Init()
	GetDB()
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "user=admin password=admin dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&author.Author{}, &Book.Book{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var hold = time.Duration(1)
		for dbase == nil {
			hold = hold * 2
			fmt.Printf("database is unavalable. Wait for  %d seconds", hold)
			time.Sleep(hold * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
