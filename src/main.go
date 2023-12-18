package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
}

func main() {

	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@127.0.0.1:5432/orm_example?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Bruno", Email: "bruno@gmail.com", Password: "123123"})
	db.Create(&User{Name: "Lucas", Email: "lucas@gmail.com", Password: "123123"})
	db.Create(&User{Name: "Jade", Email: "jade@gmail.com", Password: "123123"})

	// Read
	var user1 User
	db.First(&user1, "name = ?", "Bruno")
	fmt.Println(user1)
	db.Model(&user1).Update("Password", "123")

	var user2 User
	db.First(&user2, "name = ?", "Jade")
	fmt.Println(user2)
	db.Model(&user2).Updates(User{Email: "UpdatedEmail@gmail.com", Password: "15953"}) // non-zero fields

	var user3 User
	db.First(&user3, "name = ?", "Lucas")
	fmt.Println(user3)
	db.Model(&user3).Updates(map[string]interface{}{"Email": "Lucas@gmail.com.br", "Password": "zina"})

	// Delete - delete user
	db.Delete(&user1, "name = ?", "Bruno")
}