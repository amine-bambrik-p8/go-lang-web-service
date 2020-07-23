package services

import (
	"github.com/amine-bambrik-p8/go-lang-web-service/db"
	"github.com/amine-bambrik-p8/go-lang-web-service/models"
	"github.com/jinzhu/gorm"
)

// IUsersService represents available Users service mothods
type IUsersService interface {
	AllUsers() []models.User
	FindUser(user *models.NewUser) *models.User
	FindUserById(id uint) *models.User
	NewUser(user *models.NewUser) *models.User
	DeleteUser(id uint) *models.User
	UpdateUser(user *models.NewUser) *models.User
}

// UsersService represents available Users service struct type
type UsersService struct {
}

// Users represents Users service with the users table
var (
	Users UsersService
)

// AllUsers return the list of users
func (s *UsersService) AllUsers() []models.User {
	db, err := gorm.Open(db.Database.Dialect, db.Database.Database)
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var users []models.User
	db.Find(&users)
	return users
}

//FindUserById finds user by id and returns the user
func (s *UsersService) FindUserById(id uint64) *models.User {
	db, err := gorm.Open(db.Database.Dialect, db.Database.Database)
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var result models.User
	db.First(&result, id)
	return &result
}
func (s *UsersService) FindUser(user *models.NewUser) *models.User {
	db, err := gorm.Open(db.Database.Dialect, db.Database.Database)
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var result models.User
	db.Where("username = ?", user.Username).First(&result)
	result.CheckPasswordHash(user.Password)
	return &result
}
func (s *UsersService) NewUser(user *models.NewUser) *models.User {
	db, err := gorm.Open(db.Database.Dialect, db.Database.Database)
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	newUser := &models.User{Username: user.Username, Password: user.Password, Role: user.Role}
	db.Create(newUser)
	return newUser
}
func (s *UsersService) DeleteUser(id uint64) *models.User {
	db, err := gorm.Open(db.Database.Dialect, db.Database.Database)
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	user := s.FindUserById(id)
	db.Delete(user)
	return user
}
func (s *UsersService) UpdateUser(user *models.NewUser) *models.User {
	db, err := gorm.Open(db.Database.Dialect, db.Database.Database)
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var result models.User
	db.Where("username = ?", user.Username).First(&result)

	result.Password = user.Password
	result.Role = user.Role
	return &result
}
