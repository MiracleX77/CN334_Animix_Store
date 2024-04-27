package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/review/entities"
	reviewError "github.com/MiracleX77/CN334_Animix_Store/review/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type reviewPosgresRepository struct {
	db *gorm.DB
}

func NewReviewPostgresRepository(db *gorm.DB) ReviewRepository {
	return &reviewPosgresRepository{db: db}
}

func (r *reviewPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Review)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &reviewError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *reviewPosgresRepository) GetDataByKey(key string, value *string) (*entities.Review, error) {
	data := new(entities.Review)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &reviewError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}
func (r *reviewPosgresRepository) GetDataAll() ([]*entities.Review, error) {
	datas := []*entities.Review{}
	result := r.db.Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &reviewError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *reviewPosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Review, error) {
	datas := []*entities.Review{}
	result := r.db.Preload("Delivery").Preload("Payment").Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &reviewError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *reviewPosgresRepository) InsertData(in *entities.InsertReview) error {
	data := &entities.Review{
		UserId:    in.UserId,
		ProductId: in.ProductId,
		Title:     in.Title,
		Content:   in.Content,
		Rating:    in.Rating,
		Polarity:  in.Polarity,
		Status:    in.Status,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &reviewError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *reviewPosgresRepository) UpdateData(in *entities.UpdateReview, id *uint64) error {
	result := r.db.Model(&entities.Review{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"title":    in.Title,
		"content":  in.Content,
		"rating":   in.Rating,
		"polarity": in.Polarity,
		"status":   in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &reviewError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *reviewPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Review{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &reviewError.ServerInternalError{Err: result.Error}
	}
	return nil
}
