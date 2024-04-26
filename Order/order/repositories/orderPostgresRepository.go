package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/order/entities"
	orderError "github.com/MiracleX77/CN334_Animix_Store/order/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type orderPosgresRepository struct {
	db *gorm.DB
}

func NewOrderPostgresRepository(db *gorm.DB) OrderRepository {
	return &orderPosgresRepository{db: db}
}

func (r *orderPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Order)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &orderError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *orderPosgresRepository) GetDataByKey(key string, value *string) (*entities.Order, error) {
	data := new(entities.Order)
	data_e := r.db.Preload("Delivery").Preload("Payment").Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &orderError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}
func (r *orderPosgresRepository) GetDataAll() ([]*entities.Order, error) {
	datas := []*entities.Order{}
	result := r.db.Preload("Delivery").Preload("Payment").Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &orderError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *orderPosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Order, error) {
	datas := []*entities.Order{}
	result := r.db.Preload("Delivery").Preload("Payment").Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &orderError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *orderPosgresRepository) InsertData(in *entities.InsertOrder) (int64, error) {
	data := &entities.Order{
		UserId:     in.UserId,
		DeliveryId: in.DeliveryId,
		PaymentId:  in.PaymentId,
		TotalPrice: in.TotalPrice,
		Status:     in.Status,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return 0, &orderError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return int64(data.ID), nil
}

func (r *orderPosgresRepository) UpdateData(in *entities.UpdateOrder, id *uint64) error {
	result := r.db.Model(&entities.Order{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"status": in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &orderError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *orderPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Order{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &orderError.ServerInternalError{Err: result.Error}
	}
	return nil
}
