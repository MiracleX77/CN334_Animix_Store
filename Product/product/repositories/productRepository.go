package repositories

import "github.com/MiracleX77/CN334_Animix_Store/product/entities"

type ProductRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Product, error)
	GetDataAll() ([]*entities.Product, error)
	InsertData(in *entities.InsertProduct) error
	UpdateData(in *entities.UpdateProduct, id *uint64) error
	GetDataAllByKey(key string, value *string) ([]*entities.Product, error)
	DeleteData(id *uint64) error
}

type AuthorRepository interface {
	GetDataByKey(key string, value *string) (*entities.Author, error)
	InsertData(in *entities.InsertAuthor) error
	UpdateData(in *entities.UpdateAuthor, id *uint64) error
	GetDataAll() ([]*entities.Author, error)
	DeleteData(id *uint64) error
}

type CategoryRepository interface {
	GetDataByKey(key string, value *string) (*entities.Category, error)
	InsertData(in *entities.InsertCategory) error
	UpdateData(in *entities.UpdateCategory, id *uint64) error
	GetDataAll() ([]*entities.Category, error)
	DeleteData(id *uint64) error
}

type PublisherRepository interface {
	GetDataByKey(key string, value *string) (*entities.Publisher, error)
	InsertData(in *entities.InsertPublisher) error
	UpdateData(in *entities.UpdatePublisher, id *uint64) error
	GetDataAll() ([]*entities.Publisher, error)
	DeleteData(id *uint64) error
}
