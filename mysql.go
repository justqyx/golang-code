package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "L1212"})

	// Read
	var user User
	db.First(&user, 1)
	db.First(&user, "name = ?", "L1212")

	db.Model(&user).Update("name", "L1213")

	// Delete - deleteuser
	db.Delete(&user)
}
