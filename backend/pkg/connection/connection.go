package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"weekdemy-task-manager-backend/pkg/config"
	"weekdemy-task-manager-backend/pkg/models"
)

var db *gorm.DB

// Connect to the database.
func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIp, dbConfig.DBName)
	//root:nisaad@tcp(127.0.0.1:3306)/go-crud?charset=utf8mb4&parseTime=True&loc=Local
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	db = d
}

// migrate the tables.
func migrate() {
	err := db.Migrator().AutoMigrate(&models.BookDetail{})
	if err != nil {
		fmt.Println("Error migrating table. ", err)
		panic(err)
	}
	err = db.Migrator().AutoMigrate(&models.AuthorDetail{})
	if err != nil {
		fmt.Println("Error migrating table. ", err)
		panic(err)
	}
	err = db.Migrator().AutoMigrate(&models.UserDetail{})
	if err != nil {
		fmt.Println("Error migrating table. ", err)
		panic(err)
	}
}

// GetDB returns the db instance.
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
