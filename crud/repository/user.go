package repository

import (
	"trunne-crud/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	Repository
}

func NewUserRepository(database *gorm.DB) entity.UserRepository {
	return &userRepository{
		Repository: Repository{
			database: database,
		},
	}
}

func (ur *userRepository) Save(user entity.User) error {
	err := ur.database.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindById(id int) (*entity.User, error) {
	var user *entity.User
	err := ur.database.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := ur.database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
