package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

//users
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Fullname  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Address   string
	IsAdmin   bool `gorm:"type:bool;default:false"` // map with is_active use column:is_active
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password = hashPassword(user.Password)
	return nil
}

func hashPassword(password string) string {

	argon := argon2.DefaultConfig()
	encoded, _ := argon.HashEncoded([]byte(password))
	return string(encoded)

}
