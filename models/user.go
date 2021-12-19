package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

//users
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Fullname  string `json:"fullname" gorm:"type:varchar(100);not null"`
	Email     string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string `json:"-" gorm:"type:varchar(255);not null"`
	Address   string
	IsAdmin   bool      `json:"is_admin" gorm:"type:bool;default:false"` // map with is_active use column:is_active
	Blogs     []Blog    `json:"blogs" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
