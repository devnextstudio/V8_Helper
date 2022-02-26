package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

var (
	whost = os.Getenv("WRITE_DB_HOST")
	rhost = os.Getenv("READ_DB_HOST")
	user  = os.Getenv("DB_USER")
	name  = os.Getenv("DB_NAME")
	pass  = os.Getenv("DB_PASSWORD")
)

func WDBOpen() (db *gorm.DB) {

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		whost,
		name,
	)

	db, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return

}

func RDBOpen() (db *gorm.DB) {

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		rhost,
		name,
	)

	db, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return

}

func LowCodeWriteOpen() (db *sql.DB) {

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		whost,
		name,
	)

	db, err := sql.Open("mysql", dbSource)

	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)

	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(151)

	if err != nil {
		fmt.Println("Connection Error: ")
		fmt.Println(err.Error())
		panic(err)
	}

	return

}

func LowCodeReadOpen() (db *sql.DB) {

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		rhost,
		name,
	)

	db, err := sql.Open("mysql", dbSource)

	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)

	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(151)

	if err != nil {
		fmt.Println("Connection Error: ")
		fmt.Println(err.Error())
		panic(err)
	}

	return

}
