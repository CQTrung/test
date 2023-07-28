package usecase

import "trunne-crud/entity"

type userUsecase struct {
	userRepository entity.UserRepository
}

func NewUserUsecase(ur entity.UserRepository) entity.UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase)FindById(id int) (*entity.User, error) {
	user, err := uu.userRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uu *userUsecase)FindByEmail(email string) (*entity.User, error){
	user, err := uu.userRepository.FindByEmail(email)
}
