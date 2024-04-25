package repositories

import "github.com/MiracleX77/CN334_Animix_Store/payment/entities"

type PaymentRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Payment, error)
	InsertData(in *entities.InsertPayment) (int64, error)
}
