package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
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

/*func LowCodeOpen() (db *sql.DB) {

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		host,
		name,
	)

	db, err := sql.Open("mysql", dbSource)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	if err != nil {
		fmt.Println("Connection Error: ")
		fmt.Println(err.Error())
		panic(err)
	}

	return

}*/
