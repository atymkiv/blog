package mockdb

import (
	"fmt"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //
	"log"
)

type DB struct {
	*gorm.DB
	//CreateFn func(value interface{}) *gorm.DB
	//FindFn   func(out interface{}, where ...interface{}) *gorm.DB
	//WhereFn  func(query interface{}, args ...interface{}) *gorm.DB
}

//func (db *DB) Create(value interface{}) *gorm.DB {
//	return db.CreateFn(value)
//}
//
//func (db *DB) Find(out interface{}, where ...interface{}) *gorm.DB {
//	return db.FindFn(out, where)
//}
//
//func (db *DB) Where(query interface{}, args ...interface{}) *gorm.DB {
//	return db.WhereFn(query, args)
//}

func NewFakeDb() (*DB, error) {
	// Initialize GORM
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("Could not create the GORM database: %v", err)
	}

	// Create tables
	db.AutoMigrate(&blog.User{}, &blog.Post{})

	return &DB{db}, nil
}

func NewFakeDbOrFatal() *DB {
	db, err := NewFakeDb()
	if err != nil {
		log.Fatalf("The fake DB doesn't create successfully. Fail fast.")
	}

	return db
}
