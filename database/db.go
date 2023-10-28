package database

import (
	"github.com/rachma1201/Task-5-pbi-btpns-Rahmaliyah-Kadir/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null;size:255"`
	Photos   []Photo
}

type Photo struct {
	gorm.Model
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint
}

var DB *gorm.DB

func ConnectDB() {
	dsn := "task-5-pbi-btpns-Rahmaliyah-Kadir.database"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB = db
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}

func GetDB() *gorm.DB {
	return DB
}
