package client

import (
	"gorm.io/gorm"
)

func NewDBClient() *gorm.DB {
	/*
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			User,
			Password,
			Host,
			Port,
			DBName,
		)
	*/

	return nil
	/*
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		return db
	*/
}
