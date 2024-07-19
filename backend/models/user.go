package models

import (
	"log"
	"time"

	"app.com/db"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (e User) Save() {
	db ,_:= db.InitDB()

	
	result := db.Create(&e)
    if result.Error != nil {
        log.Fatalf("failed to create user: %v", result.Error)
    }


}
// func GetAllEvent() []User {
// 	return user
// }
