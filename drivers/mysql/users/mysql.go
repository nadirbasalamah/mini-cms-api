package users

import (
	"mini-cms-api/businesses/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(userReq *users.Domain) (users.Domain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return users.Domain{}, err
	}

	record := FromDomain(userReq)

	record.Password = string(password)

	result := ur.conn.Create(&record)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return users.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *userRepository) GetUserByEmail(userReq *users.Domain) (users.Domain, error) {
	var user User

	err := ur.conn.First(&user, "email = ?", userReq.Email).Error

	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (ur *userRepository) GetUserInfo(id string) (users.Domain, error) {
	var user User

	err := ur.conn.First(&user, "id = ?", id).Error

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}
