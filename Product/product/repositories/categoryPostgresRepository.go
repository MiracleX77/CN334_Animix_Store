package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type categoryPosgresRepository struct {
	db *gorm.DB
}

func NewCategoryPostgresRepository(db *gorm.DB) CategoryRepository {
	return &categoryPosgresRepository{db: db}
}

func (r *categoryPosgresRepository) GetDataByKey(key string, value *string) (*entities.Category, error) {
	data := new(entities.Category)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &productError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}

func (r *categoryPosgresRepository) GetDataAll() ([]*entities.Category, error) {
	datas := []*entities.Category{}
	result := r.db.Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *categoryPosgresRepository) InsertData(in *entities.InsertCategory) error {
	data := &entities.Category{
		Name:   in.Name,
		Status: "active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *categoryPosgresRepository) UpdateData(in *entities.UpdateCategory, id *uint64) error {
	result := r.db.Model(&entities.Category{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"name":   in.Name,
		"status": in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *categoryPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Category{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	return nil
}
