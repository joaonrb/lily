//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package auth

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserHash string `gorm:"not null;unique"`
	Name     string `gorm:"not null;unique"`
	Password string
	IsAdmin  bool
}

type AuthToken struct {
	ID        uint
	UserID    User
	Token     string `gorm:"not null;unique"`
	CreatedAt time.Time
	Duration  uint
	IsDeleted bool
}

func RegisterModels(db gorm.DB) {
	db.AutoMigrate(&User{}, &AuthToken{})
}