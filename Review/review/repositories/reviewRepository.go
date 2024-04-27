package repositories

import "github.com/MiracleX77/CN334_Animix_Store/review/entities"

type ReviewRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Review, error)
	GetDataAll() ([]*entities.Review, error)
	InsertData(in *entities.InsertReview) error
	UpdateData(in *entities.UpdateReview, id *uint64) error
	GetDataAllByKey(key string, value *string) ([]*entities.Review, error)
	DeleteData(id *uint64) error
}
