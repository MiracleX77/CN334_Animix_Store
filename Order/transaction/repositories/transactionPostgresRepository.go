package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/transaction/entities"
	transactionError "github.com/MiracleX77/CN334_Animix_Store/transaction/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type transactionPosgresRepository struct {
	db *gorm.DB
}

func NewTransactionPostgresRepository(db *gorm.DB) TransactionRepository {
	return &transactionPosgresRepository{db: db}
}

func (r *transactionPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Transaction)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &transactionError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *transactionPosgresRepository) GetDataByKey(key string, value *string) (*entities.Transaction, error) {
	data := new(entities.Transaction)
	data_e := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &transactionError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}
func (r *transactionPosgresRepository) GetDataAll() ([]*entities.Transaction, error) {
	datas := []*entities.Transaction{}
	result := r.db.Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &transactionError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *transactionPosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Transaction, error) {
	datas := []*entities.Transaction{}
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &transactionError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *transactionPosgresRepository) InsertData(in *entities.InsertTransaction) (int64, error) {
	data := &entities.Transaction{
		ProductId: in.ProductId,
		OrderId:   in.OrderId,
		Status:    "Active",
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return 0, &transactionError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return result.RowsAffected, nil
}

func (r *transactionPosgresRepository) UpdateData(in *entities.UpdateTransaction, id *uint64) error {
	result := r.db.Model(&entities.Transaction{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"product_id": in.ProductId,
		"order_id":   in.OrderId,
		"status":     in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &transactionError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *transactionPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Transaction{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &transactionError.ServerInternalError{Err: result.Error}
	}
	return nil
}
