package repositories

import "github.com/MiracleX77/CN334_Animix_Store/transaction/entities"

type TransactionRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Transaction, error)
	GetDataAll() ([]*entities.Transaction, error)
	InsertData(in *entities.InsertTransaction) (int64, error)
	UpdateData(in *entities.UpdateTransaction, id *uint64) error
	GetDataAllByKey(key string, value *string) ([]*entities.Transaction, error)
	DeleteData(id *uint64) error
}
