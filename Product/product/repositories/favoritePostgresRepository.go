package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type favoritePosgresRepository struct {
	db *gorm.DB
}

func NewFavoritePostgresRepository(db *gorm.DB) FavoriteRepository {
	return &favoritePosgresRepository{db: db}
}

func (r *favoritePosgresRepository) GetDataByKey(key string, value *string) (*entities.Favorite, error) {
	data := new(entities.Favorite)
	data_e := r.db.Preload("Product").Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &productError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}

func (r *favoritePosgresRepository) GetDataAll() ([]*entities.Favorite, error) {
	datas := []*entities.Favorite{}
	result := r.db.Preload("Product").Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *favoritePosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Favorite, error) {
	datas := []*entities.Favorite{}
	result := r.db.Preload("Product").Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *favoritePosgresRepository) InsertData(in *entities.InsertFavorite) error {
	data := &entities.Favorite{
		ProductId: in.ProductId,
		UserId:    in.UserId,
		Status:    "Active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *favoritePosgresRepository) UpdateData(in *entities.UpdateFavorite, id *uint64) error {
	result := r.db.Model(&entities.Favorite{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"status": in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *favoritePosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Favorite{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	return nil
}
