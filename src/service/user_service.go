package service

import (
	"oauth2/src/common"
	"oauth2/src/models"
	"oauth2/src/repository"
)

type UserService interface {
	LoginByUserAndPassword(username string, password string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(uuid string) error
	FindById(id string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error)
}

type userRepositoryImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userRepositoryImpl {
	return &userRepositoryImpl{userRepository}
}

func (u *userRepositoryImpl) LoginByUserAndPassword(username string, password string) (*models.User, error) {
	hash := common.NewHash()
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if result, err := hash.BCryptCompare(*user.Password, password); result {
		return  user, err
	}

	return nil, err
}


func (u *userRepositoryImpl) Create(user *models.User) error {
	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImpl) Update(user *models.User) error {
	err := u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImpl) Delete(uuid string) (error) {
	err := u.userRepository.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImpl) FindById(uuid string) (*models.User, error){
	user, err := u.userRepository.FindById(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImpl) Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error) {
	users, err := u.userRepository.Paginate(filters, orderBy, orderDir, limit, page)
	if err != nil {
		return nil, err
	}
	return users, nil
}
