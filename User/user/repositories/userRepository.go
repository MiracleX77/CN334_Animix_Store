package repositories

import "github.com/MiracleX77/CN334_Animix_Store/user/entities"

type UserRepository interface {
	Search(key string, value *string) (bool, error)
	GetUserDataByKey(key string, value *string) (*entities.User, error)
	InsertUserData(in *entities.InsertUser) error
	UpdateUserData(in *entities.UpdateUser, id *uint64) error
	GetUserDataAll() ([]*entities.User, error)
	DeleteUserData(id *uint64) error
}
