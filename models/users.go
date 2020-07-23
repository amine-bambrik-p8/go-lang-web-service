package models

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// GetViewModel get Executed before the object get sent to the client
// Returns The User's View Model
func (u *User) GetViewModel() interface{} {
	return map[string]interface{}{
		"username": u.Username,
		"role":     u.Role,
	}
}

// CheckPasswordHash checks whether the given password is correct or not
// Returns true if password matchs else false
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// hashPassword takes a Password and returns a hash of the password
// Return the hashed password or err in case of failure
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// BeforeSave Hash the password before saving to the database
func (u *User) BeforeSave() (err error) {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hash
	return
}

// GenerateJWT generates JWT token to be used for auth
func (u *User) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["role"] = u.Role

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	return tokenString, nil
}

// BeforeSave Hash the password before update to the database
// NOTE this function will be called by the GORM before update
func (u *User) BeforeUpdate() (err error) {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hash
	return
}

// NewUser represents the user sent by the client as http request body
// TODO should be renamed
type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// User represents the model of the users table also includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt.
// Note Table name is users
type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique;not null"`
	Role     string `gorm:"column:role,default:USER"`
	Password string `gorm:"column:password;not null"`
}
