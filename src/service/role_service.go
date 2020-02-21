package service

import (
	"oauth2/src/common"
	"oauth2/src/models"
	"oauth2/src/repository"
)

type RoleService interface {
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(uuid string) error
	FindById(uuid string) (*models.Role, error)
	FindByName(username string) (*models.Role, error)
	Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error)
}

type roleRepositoryImpl struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *roleRepositoryImpl {
	return &roleRepositoryImpl{roleRepository}
}


func (u *roleRepositoryImpl) Create(role *models.Role) error {
	err := u.roleRepository.Create(role)
	if err != nil {
		return err
	}
	return nil
}

func (u *roleRepositoryImpl) Update(role *models.Role) error {
	err := u.roleRepository.Update(role)
	if err != nil {
		return err
	}
	return nil
}

func (u *roleRepositoryImpl) Delete(uuid string) error {
	err := u.roleRepository.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (u *roleRepositoryImpl) FindById(uuid string) (*models.Role, error){
	user, err := u.roleRepository.FindById(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *roleRepositoryImpl) FindByName(name string) (*models.Role, error) {
	user, err := u.roleRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *roleRepositoryImpl) Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error) {
	users, err := u.roleRepository.Paginate(filters, orderBy, orderDir, limit, page)
	if err != nil {
		return nil, err
	}
	return users, nil
}
