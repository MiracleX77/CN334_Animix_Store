package usecases

import (
	"errors"

	authError "github.com/MiracleX77/CN334_Animix_Store/auth/errors"
	"github.com/MiracleX77/CN334_Animix_Store/auth/models"
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	"github.com/MiracleX77/CN334_Animix_Store/user/entities"
	"github.com/MiracleX77/CN334_Animix_Store/user/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	RegisterDataProcessing(in *models.RegisterData) error
	CheckData(in *models.RegisterData) error
	LoginDataProcession(in *models.LoginData) (*string, error)
}

type authUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewAuthUsecaseImpl(userRepository repositories.UserRepository) AuthUsecase {
	return &authUsecaseImpl{
		userRepository: userRepository,
	}
}
func (u *authUsecaseImpl) CheckData(in *models.RegisterData) error {
	username := &in.Username
	//result = true -> found username
	//result = false -> Not found username
	//err != nill -> Found error
	if result, err := u.userRepository.Search("username", username); result || err != nil {
		return &authError.UsernameAlreadyExistError{}
	} else {
		if errors.Is(err, &authError.ServerInternalError{}) {
			return &authError.ServerInternalError{Err: err}
		}
	}
	return nil
}

func (u *authUsecaseImpl) RegisterDataProcessing(in *models.RegisterData) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &authError.ServerInternalError{Err: err}
	}
	insertUserData := &entities.InsertUser{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Username:  in.Username,
		Password:  string(hashedPassword),
		Type:      "User",
		Status:    "Active",
	}

	if err := u.userRepository.InsertUserData(insertUserData); err != nil {
		return &authError.ServerInternalError{Err: err}
	}

	return nil
}

func (u *authUsecaseImpl) LoginDataProcession(in *models.LoginData) (*string, error) {
	username := &in.Username
	password := &in.Password
	//result = true -> found username
	//result = false -> Not found username
	//err != nill -> Found error
	if result, err := u.userRepository.Search("username", username); !result || err != nil {
		if err != nil {
			return nil, &authError.ServerInternalError{Err: err}
		}
		return nil, &authError.UsernameNotFoundError{}
	}
	if user, err := u.userRepository.GetUserDataByKey("username", username); err != nil {
		//return error
		return nil, &authError.ServerInternalError{Err: err}
	} else {
		//compare password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*password)); err != nil {
			return nil, &authError.PasswordIncorrectError{}
		} else {
			//return success
			//generate token
			tokenUsecase := NewTokenUsecaseImpl(configs.GetJwtConfig().SecretKey)
			token, err := tokenUsecase.GenerateToken(&user.ID, &user.Username, &user.Type)
			if err != nil {
				return nil, &authError.ServerInternalError{Err: err}
			}
			return token, nil
		}
	}
}
