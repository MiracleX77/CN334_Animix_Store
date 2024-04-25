package repositories

import (
	"github.com/MiracleX77/CN334_Animix_Store/auth/entities"
	userError "github.com/MiracleX77/CN334_Animix_Store/auth/errors"

	"gorm.io/gorm"
)

type userPosgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserRepository {
	return &userPosgresRepository{db: db}
}

func (r *userPosgresRepository) Search(key string, value *string) (bool, error) {
	user := new(entities.User)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(user)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &userError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *userPosgresRepository) GetUserDataByKey(key string, value *string) (*entities.User, error) {
	user := new(entities.User)
	user_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(user)
	if user_data.Error != nil {
		return nil, &userError.ServerInternalError{Err: user_data.Error}
	}
	return user, nil
}
