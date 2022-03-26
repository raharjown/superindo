package Models

import (
	"errors"
	"superindo/Database"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Password string    `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func Login(userLogin *User) (err error) {
	var user User
	err = Database.DB.Where("email=?", userLogin.Email).First(&user).Error
	if err != nil {
		return err
	}
	if user.Password != userLogin.Password {
		return errors.New("Wrong Password")
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	err = Database.DB.Where("id=?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = Database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
