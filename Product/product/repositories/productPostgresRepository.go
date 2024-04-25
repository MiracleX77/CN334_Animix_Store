package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type productPosgresRepository struct {
	db *gorm.DB
}

func NewProductPostgresRepository(db *gorm.DB) ProductRepository {
	return &productPosgresRepository{db: db}
}

func (r *productPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Product)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &productError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *productPosgresRepository) GetDataByKey(key string, value *string) (*entities.Product, error) {
	data := new(entities.Product)
	if key == "name" && value != nil {
		key = key + " LIKE ?"
		*value = "%" + *value + "%" // This modifies the search term to be a wildcard search
	} else if value != nil {
		key = key + " = ?"
	}
	data_e := r.db.Preload("Author").Preload("Publisher").Preload("Category").Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &productError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}
func (r *productPosgresRepository) GetDataAll() ([]*entities.Product, error) {
	datas := []*entities.Product{}
	result := r.db.Preload("Author").Preload("Publisher").Preload("Category").Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *productPosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Product, error) {
	datas := []*entities.Product{}
	result := r.db.Preload("Author").Preload("Publisher").Preload("Category").Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &productError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *productPosgresRepository) InsertData(in *entities.InsertProduct) error {
	data := &entities.Product{
		AuthorId:    in.AuthorId,
		PublisherId: in.PublisherId,
		CategoryId:  in.CategoryId,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Stock:       in.Stock,
		ImgUrl:      in.ImgUrl,
		Status:      "Active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *productPosgresRepository) UpdateData(in *entities.UpdateProduct, id *uint64) error {
	result := r.db.Model(&entities.Product{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"author_id":    in.AuthorId,
		"publisher_id": in.PublisherId,
		"category_id":  in.CategoryId,
		"name":         in.Name,
		"description":  in.Description,
		"price":        in.Price,
		"stock":        in.Stock,
		"img_url":      in.ImgUrl,
		"status":       in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *productPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Product{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &productError.ServerInternalError{Err: result.Error}
	}
	return nil
}
