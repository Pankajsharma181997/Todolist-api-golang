package connectors

import (
	"example/todolist-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {

	dsn := "root:Pankaj_18@tcp(127.0.0.1:4545)/todos?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(" Failed to connect to DB")
	}
	// Migrate
	db.AutoMigrate(&models.Todo{})
	return db
}

func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection to the database")
	}
	dbSQL.Close()
}
