package config

import (
	"fmt"
	"praktikum/rest/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	config := model.Config{
		DBUsername: "root",
		DBPassword: "",
		DBPort:     "3306",
		DBHost:     "localhost",
		DBName:     "orm_golang",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}


func InitialMigration() {
	DB.AutoMigrate(&model.User{}, &model.Books{})
}

func InitDBTest() {
	const DB_USER_TEST = "root"
	const DB_PASS_TEST = ""
	const DB_HOST_TEST = "localhost"
	const DB_PORT_TEST = "3306"
	const DB_NAME_TEST = "go_test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&model.User{})
	DB.AutoMigrate(&model.User{})
	DB.Migrator().DropTable(&model.Books{})
	DB.AutoMigrate(&model.Books{})
}


func InitDBTestNoDrop() {
	const DB_USER_TEST = "root"
	const DB_PASS_TEST = ""
	const DB_HOST_TEST = "localhost"
	const DB_PORT_TEST = "3306"
	const DB_NAME_TEST = "go_test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTestNoDrop()
}


func InitMigrateTestNoDrop() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Books{})
}

