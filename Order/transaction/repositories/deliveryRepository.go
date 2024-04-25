package repositories

import "github.com/MiracleX77/CN334_Animix_Store/delivery/entities"

type DeliveryRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Delivery, error)
	GetDataAll() ([]*entities.Delivery, error)
	InsertData(in *entities.InsertDelivery) (int64, error)
	UpdateData(in *entities.UpdateDelivery, id *uint64) error
	GetDataAllByKey(key string, value *string) ([]*entities.Delivery, error)
	DeleteData(id *uint64) error
}
