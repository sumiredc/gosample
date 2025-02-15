package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, *gorm.DB, error) {
	writerDNS := "user:password@tcp(mysql:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	writer, err := gorm.Open(mysql.Open(writerDNS), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	readerDNS := "user:password@tcp(mysql:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	reader, err := gorm.Open(mysql.Open(readerDNS), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return writer, reader, nil
}
