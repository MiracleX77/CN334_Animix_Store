package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type authorPosgresRepository struct {
	db *gorm.DB
}

func NewAuthorPostgresRepository(db *gorm.DB) AuthorRepository {
	return &authorPosgresRepository{db: db}
}

func (r *authorPosgresRepository) GetDataByKey(key string, value *string) (*entities.Author, error) {
	data := new(entities.Author)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &productError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}

func (r *authorPosgresRepository) GetDataAll() ([]*entities.Author, error) {
	datas := []*entities.Author{}
	result := r.db.Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *authorPosgresRepository) InsertData(in *entities.InsertAuthor) error {
	data := &entities.Author{
		Name:        in.Name,
		Description: in.Description,
		Status:      "active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *authorPosgresRepository) UpdateData(in *entities.UpdateAuthor, id *uint64) error {
	result := r.db.Model(&entities.Author{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"name":        in.Name,
		"description": in.Description,
		"status":      in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *authorPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Author{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	return nil
}
