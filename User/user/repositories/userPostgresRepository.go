package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/user/entities"
	userError "github.com/MiracleX77/CN334_Animix_Store/user/errors"

	"github.com/labstack/gommon/log"

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

func (r *userPosgresRepository) GetUserDataAll() ([]*entities.User, error) {
	users := []*entities.User{}
	result := r.db.Where("status <> ?", "Removed").Find(&users)
	if result.Error != nil {
		return nil, &userError.ServerInternalError{Err: result.Error}
	}
	return users, nil
}

func (r *userPosgresRepository) InsertUserData(in *entities.InsertUser) error {
	data := &entities.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  in.Password,
		Email:     in.Email,
		Type:      in.Type,
		Status:    in.Status,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertUserData:%v", result.Error)
		return &userError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertUserData: %v", result.RowsAffected)
	return nil
}

func (r *userPosgresRepository) UpdateUserData(in *entities.UpdateUser, id *uint64) error {
	result := r.db.Model(&entities.User{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"first_name": in.FirstName,
		"last_name":  in.LastName,
		"email":      in.Email,
		"status":     "Active",
	})
	if result.Error != nil {
		log.Errorf("UpdateUserData:%v", result.Error)
		return &userError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *userPosgresRepository) DeleteUserData(id *uint64) error {
	result := r.db.Model(&entities.User{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteUserData:%v", result.Error)
		return &userError.ServerInternalError{Err: result.Error}
	}
	return nil
}
