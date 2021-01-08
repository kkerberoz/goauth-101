package models

import (
	"ginauth101/utils/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goonode/mogo"
)

//User struct is to handle user data
type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Email              string `idx:"{email},unique" json:"email" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Name               string `json:"name"`
	// CreatedAt          *time.Time
	// UpdatedAt          *time.Time
	VerifiedAt *time.Time
}

//GetJwtToken returns jwt token with user email claims
func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
	})
	secretKey := config.EnvVar("TOKEN_KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func init() {
	mogo.ModelRegistry.Register(User{})
}
