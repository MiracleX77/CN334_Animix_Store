package repositories

import (
	"github.com/MiracleX77/CN334_Animix_Store/payment/entities"
	paymentError "github.com/MiracleX77/CN334_Animix_Store/payment/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type paymentPosgresRepository struct {
	db *gorm.DB
}

func NewPaymentPostgresRepository(db *gorm.DB) PaymentRepository {
	return &paymentPosgresRepository{db: db}
}

func (r *paymentPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Payment)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &paymentError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *paymentPosgresRepository) GetDataByKey(key string, value *string) (*entities.Payment, error) {
	data := new(entities.Payment)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &paymentError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}

func (r *paymentPosgresRepository) InsertData(in *entities.InsertPayment) (int64, error) {
	data := &entities.Payment{
		Type:         in.Type,
		Total:        in.Total,
		ProofPayment: in.ProofPayment,
		Status:       "Active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return 0, &paymentError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return result.RowsAffected, nil
}
