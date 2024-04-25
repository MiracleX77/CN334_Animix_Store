package repositories

import (
	"time"

	"github.com/MiracleX77/CN334_Animix_Store/address/entities"
	addressError "github.com/MiracleX77/CN334_Animix_Store/address/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type addressPosgresRepository struct {
	db *gorm.DB
}

func NewAddressPostgresRepository(db *gorm.DB) AddressRepository {
	return &addressPosgresRepository{db: db}
}

func (r *addressPosgresRepository) Search(key string, value *string) (bool, error) {
	data := new(entities.Address)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(data)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &addressError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *addressPosgresRepository) GetDataByKey(key string, value *string) (*entities.Address, error) {
	data := new(entities.Address)
	data_e := r.db.Preload("SubDistrict").Preload("District").Preload("Province").Where(key+"= ?", *value).Where("status <> ?", "Removed").First(data)
	if data_e.Error != nil {
		return nil, &addressError.ServerInternalError{Err: data_e.Error}
	}
	return data, nil
}

func (r *addressPosgresRepository) GetDataAllByKey(key string, value *string) ([]*entities.Address, error) {
	datas := []*entities.Address{}
	result := r.db.Preload("SubDistrict").Preload("District").Preload("Province").Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&datas)
	if result.Error != nil {
		return nil, &addressError.ServerInternalError{Err: result.Error}
	}
	return datas, nil
}

func (r *addressPosgresRepository) InsertData(in *entities.InsertAddress) error {
	data := &entities.Address{
		UserId:        in.UserId,
		AddressLine:   in.AddressLine,
		Phone:         in.Phone,
		Name:          in.Name,
		SubDistrictId: in.SubDistrictId,
		DistrictId:    in.DistrictId,
		ProvinceId:    in.ProvinceId,
		Default:       in.Default,
		Status:        in.Status,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertData:%v", result.Error)
		return &addressError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertData: %v", result.RowsAffected)
	return nil
}

func (r *addressPosgresRepository) UpdateData(in *entities.UpdateAddress, id *uint64) error {
	result := r.db.Model(&entities.Address{}).Where("id = ?", *id).Updates(map[string]interface{}{
		"address_line":    in.AddressLine,
		"phone":           in.Phone,
		"name":            in.Name,
		"sub_district_id": in.SubDistrictId,
		"district_id":     in.DistrictId,
		"province_id":     in.ProvinceId,
		"default":         in.Default,
		"status":          in.Status,
	})
	if result.Error != nil {
		log.Errorf("UpdateData:%v", result.Error)
		return &addressError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdateUserData: %v", result.RowsAffected)
	return nil
}

func (r *addressPosgresRepository) DeleteData(id *uint64) error {
	result := r.db.Model(&entities.Address{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteData:%v", result.Error)
		return &addressError.ServerInternalError{Err: result.Error}
	}
	return nil
}

func (r *addressPosgresRepository) GetProvince() ([]*entities.Province, error) {
	provinces := []*entities.Province{}
	result := r.db.Find(&provinces)
	if result.Error != nil {
		return nil, &addressError.ServerInternalError{Err: result.Error}
	}
	return provinces, nil
}

func (r *addressPosgresRepository) GetDistrictByProvinceId(provinceId *string) ([]*entities.District, error) {
	districts := []*entities.District{}
	result := r.db.Where("province_id = ?", *provinceId).Find(&districts)
	if result.Error != nil {
		return nil, &addressError.ServerInternalError{Err: result.Error}
	}
	return districts, nil
}

func (r *addressPosgresRepository) GetSubDistrictByDistrictId(districtId *string) ([]*entities.SubDistrict, error) {
	subDistricts := []*entities.SubDistrict{}
	result := r.db.Where("district_id = ?", *districtId).Find(&subDistricts)
	if result.Error != nil {
		return nil, &addressError.ServerInternalError{Err: result.Error}
	}
	return subDistricts, nil
}
