package driver

import (
	"fmt"
	"simpledy/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var user = model.User{
// 	Username: "testUser",
// 	Password: "123456",
// }

// var userinformation = model.UserInformation{
// 	Id:            1,
// 	Name:          user.Username,
// 	FollowCount:   5050,
// 	FollowerCount: 2040,
// 	// IsFollow:      false,
// }

// var video = model.Video{
// 	Id:            2333,
// 	Author:        userinformation,
// 	PlayUrl:       "http://www.baidu.com",
// 	CoverUrl:      "http://www.google.com",
// 	FavoriteCount: 50,
// 	CommentCount:  12,
// 	IsFavorite:    true,
// }

func InitDB() (*gorm.DB, error) {
	fmt.Println("开始连接数据库")
	// 建立数据库连接
	host := "127.0.0.1"
	port := "3306"
	DBName := "simpledy"
	username := "root"
	password := ""
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		DBName,
		charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接数据库成功............")

	//建立User表
	m1 := db.Migrator()
	err1 := m1.CreateTable(&model.User{})
	if err1 != nil {
		fmt.Println("建立User表失败")
	} else {
		fmt.Println("建立User表成功......")
	}
	// db.Create(&user)

	//建立UserInformation表
	m2 := db.Migrator()
	err2 := m2.CreateTable(&model.UserInformation{})
	if err2 != nil {
		fmt.Println("建立UserInformation表失败")
	} else {
		fmt.Println("建立UserInformation表成功......")
	}
	// db.Create(&userinformation)

	//建立Video表
	m3 := db.Migrator()
	err3 := m3.CreateTable(&model.Video{})
	if err3 != nil {
		fmt.Println("建立Video表失败")
	} else {
		fmt.Println("建立Video表成功......")
	}
	// db.Create(&video)

	// m4 := db.Migrator()
	// err4 := m4.CreateTable(&model.Comment{})
	// if err4 != nil {
	// 	fmt.Println("建立Comment表失败")
	// } else {
	// 	fmt.Println("建立Comment表成功......")
	// }
	// db.Create(&comment)

	return db, err
}

// func InitData() {
// 	// 数据库自动建表
// 	db, _ := InitDB()
// 	// db.AutoMigrate(&model.User{})
// 	// db.Create(&user)
// 	m := db.Migrator()
// 	err := m.CreateTable(&model.User{})
// 	if err != nil {
// 		fmt.Println("sssssssssssssssssssssssssssssssssssssssssss")
// 	} else {
// 		fmt.Println("建表成功......")
// 	}
// 	fmt.Println("44444444444444444444444444444444")
// 	db.AutoMigrate(&model.Video{})
// 	db.Create(&video)
// }
