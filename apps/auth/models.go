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

const (
	GET_USER_FROM_AUTH_TOKEN =
		`SELECT %s
		 FROM
             users U
             LEFT OUTER JOIN auth_tocken A ON A.user_id = U.id
         WHERE A.token = "%s"
         `
)

// User query attributes
const (
	USER_ID         = "U.id"
	USER_HASH       = "U.user_hash"
	USER_EMAIL      = "U.email"
	USER_PASSWORD   = "U.password"
	USER_IS_ADMIN   = "U.is_admin"
	USER_IS_STAFF   = "U.is_staff"
	USER_CREATED_AT = "U.created_at"
	USER_UPDATED_AT = "U.updated_at"
	USER_DELETED_AT = "U.deleted_at"
)

// Auth query attributes
const (
	AUTH_ID         = "I.id"
	AUTH_USER_ID    = "I.user_id"
	AUTH_TOKEN      = "I.token"
	AUTH_SESSION    = "I.active_session"
	AUTH_CREATED_AT = "I.created_at"
	AUTH_DURATION   = "I.duration"

)

// Commented fields are implicit by gorm.Model
type User struct {
	gorm.Model
	//ID       uint       `gorm:"primary_key"`
	Hash      string    `gorm:"not null;unique;index"`
	Email     string    `gorm:"not null;unique;index"`
	Password  string
	IsAdmin   bool
	IsStaff   bool
	Token     AuthToken
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt *time.Time
}

// Defines the name equal to the default name.
// Just for explict purpose.
func (User) TableName() string {
	return "users"
}

type AuthToken struct {
	ID        uint   `gorm:"primary_key"`
	UserId    uint
	Token     string `gorm:"not null;unique;index"`
	Session   string `gorm:"not null;unique"`
	CreatedAt time.Time
	Duration  uint
}

func (AuthToken) TableName() string {
	return "auth_token"
}



func RegisterModels(db gorm.DB) {
	db.AutoMigrate(&User{}, &AuthToken{})
}