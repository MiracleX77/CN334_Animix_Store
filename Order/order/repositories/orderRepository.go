package repositories

import "github.com/MiracleX77/CN334_Animix_Store/order/entities"

type OrderRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Order, error)
	GetDataAll() ([]*entities.Order, error)
	InsertData(in *entities.InsertOrder) (int64, error)
	UpdateData(in *entities.UpdateOrder, id *uint64) error
	GetDataAllByKey(key string, value *string) ([]*entities.Order, error)
	DeleteData(id *uint64) error
}
