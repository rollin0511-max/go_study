package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study/gorm/dbpractice"
)

func main() {
	// 使用gorm 链接mysql数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/gostudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	dbpractice.Run(db)

}
