package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/delivery/entities"
	deliveryError "github.com/MiracleX77/CN334_Animix_Store/delivery/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type deliveryPosgresRepository struct {
	db *gorm.DB
}

func NewDeliveryPostgresRepository(db *gorm.DB) DeliveryRepository {
	return &deliveryPosgresRepository{db: db}
}

func (r *deliveryPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Delivery)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &deliveryError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *deliveryPosgresRepository) GetDataByKey(key string, value *string) (*entities.Delivery, error) {
	data := new(entities.Delivery)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &deliveryError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}
func (r *deliveryPosgresRepository) GetDataAll() ([]*entities.Delivery, error) {
	datas := []*entities.Delivery{}
	result := r.db.Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &deliveryError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *deliveryPosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Delivery, error) {
	datas := []*entities.Delivery{}
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &deliveryError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *deliveryPosgresRepository) InsertData(in *entities.InsertDelivery) (int64, error) {
	data := &entities.Delivery{
		AddressId: in.AddressId,
		Status:    in.Status,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return 0, &deliveryError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return int64(data.ID), nil
}

func (r *deliveryPosgresRepository) UpdateData(in *entities.UpdateDelivery, id *uint64) error {
	result := r.db.Model(&entities.Delivery{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"address_id":      in.AddressId,
		"cost":            in.Cost,
		"type":            in.Type,
		"tracking_number": in.TrackingNumber,
		"status":          in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &deliveryError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *deliveryPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Delivery{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &deliveryError.ServerInternalError{Err: result.Error}
	}
	return nil
}
