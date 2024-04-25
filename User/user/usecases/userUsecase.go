package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/user/entities"
	userError "github.com/MiracleX77/CN334_Animix_Store/user/errors"
	"github.com/MiracleX77/CN334_Animix_Store/user/models"
	"github.com/MiracleX77/CN334_Animix_Store/user/repositories"
)

type UserUsecase interface {
	GetUserById(id *string) (*models.UserModel, error)
	UpdateUser(in *models.UpdateModel, id *string) error
	CheckUserId(id *string) error
	GetUserAll() ([]*models.UserModel, error)
	DeleteUser(id *string) error
}

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewUserUsecaseImpl(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *userUsecaseImpl) CheckUserId(id *string) error {
	if result, err := u.userRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &userError.ServerInternalError{Err: err}
		}
		return &userError.UserNotFoundError{}
	}
	return nil
}

func (u *userUsecaseImpl) GetUserById(id *string) (*models.UserModel, error) {
	userData, err := u.userRepository.GetUserDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	userModel := &models.UserModel{
		ID:        uint64(userData.ID),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Username:  userData.Username,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		Status:    userData.Status,
	}
	return userModel, nil
}

func (u *userUsecaseImpl) GetUserAll() ([]*models.UserModel, error) {
	users, err := u.userRepository.GetUserDataAll()
	if err != nil {
		return nil, err
	}
	userModels := []*models.UserModel{}
	for _, user := range users {
		userModel := &models.UserModel{
			ID:        uint64(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Status:    user.Status,
		}
		userModels = append(userModels, userModel)
	}
	return userModels, nil
}

func (u *userUsecaseImpl) UpdateUser(in *models.UpdateModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &userError.ServerInternalError{Err: err}
	}
	updateUserData := &entities.UpdateUser{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Status:    "Active",
	}
	if err := u.userRepository.UpdateUserData(updateUserData, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *userUsecaseImpl) DeleteUser(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &userError.ServerInternalError{Err: err}
	}
	if err := u.userRepository.DeleteUserData(&idUint64); err != nil {
		return err
	}
	return nil
}
