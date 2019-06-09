package gorm

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/config"
	"github.com/cenkalti/backoff"
	_ "github.com/go-sql-driver/mysql" //
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func New(cfg *config.Database) (*gorm.DB, error) {
	//connectString := fmt.Sprintf("%s:%s@mysql/%s?charset=utf8&parseTime=True&loc=Local", username, password, name)

	/*conf := mysql.Config{
		User:   cfg.Name,
		DBName: cfg.Name,
		Passwd: cfg.Password,
		Addr:   cfg.Host + cfg.Port,
		Net:    "tcp",
	}*/

	//fmt.Println(conf.FormatDSN())
	connectString := "root:example@tcp(mysql:3306)/blog?charset=utf8&parseTime=True&loc=Local"
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 5 * time.Second

	var db gorm.DB
	operation := func() error {
		dbT, err := gorm.Open("mysql", connectString)
		if err != nil {
			return err
		}
		db = *dbT
		return nil
	}

	err := backoff.Retry(operation, b)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
	//defer db.Close()

	db.AutoMigrate(&blog.User{}, &blog.Post{})

	return &db, nil

	//if err != nil {
	//	log.Println("Connecting to mysql failed, trying again in 1 second")
	//	time.Sleep(1 * time.Second)
	//	db, err = gorm.Open("mysql", connectString)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

}
