package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type publisherPosgresRepository struct {
	db *gorm.DB
}

func NewPublisherPostgresRepository(db *gorm.DB) PublisherRepository {
	return &publisherPosgresRepository{db: db}
}

func (r *publisherPosgresRepository) GetDataByKey(key string, value *string) (*entities.Publisher, error) {
	data := new(entities.Publisher)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &productError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}

func (r *publisherPosgresRepository) GetDataAll() ([]*entities.Publisher, error) {
	datas := []*entities.Publisher{}
	result := r.db.Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *publisherPosgresRepository) InsertData(in *entities.InsertPublisher) error {
	data := &entities.Publisher{
		Name:   in.Name,
		Status: "Active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *publisherPosgresRepository) UpdateData(in *entities.UpdatePublisher, id *uint64) error {
	result := r.db.Model(&entities.Publisher{}).Where("id = ?", *id).Updates(map[string]interface{}{
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

func (r *publisherPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Publisher{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	return nil
}
