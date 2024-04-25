package repositories

import "github.com/MiracleX77/CN334_Animix_Store/address/entities"

type AddressRepository interface {
	Search(key string, value *string) (bool, error)
	GetDataByKey(key string, value *string) (*entities.Address, error)
	InsertData(in *entities.InsertAddress) error
	UpdateData(in *entities.UpdateAddress, id *uint64) error
	GetDataAllByKey(key string, value *string) ([]*entities.Address, error)
	DeleteData(id *uint64) error
	GetProvince() ([]*entities.Province, error)
	GetDistrictByProvinceId(provinceId *string) ([]*entities.District, error)
	GetSubDistrictByDistrictId(districtId *string) ([]*entities.SubDistrict, error)
}
