package repositories

import "github.com/MiracleX77/CN334_Animix_Store/auth/entities"

type UserRepository interface {
	Search(key string, value *string) (bool, error)
	GetUserDataByKey(key string, value *string) (*entities.User, error)
}
