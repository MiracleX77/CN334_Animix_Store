package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/address/entities"
	addressError "github.com/MiracleX77/CN334_Animix_Store/address/errors"
	"github.com/MiracleX77/CN334_Animix_Store/address/models"
	"github.com/MiracleX77/CN334_Animix_Store/address/repositories"
)

type AddressUsecase interface {
	InsertAddress(in *models.InsertAddressModel, id *string) error
	GetAddressById(id *string) (*models.AddressModel, error)
	UpdateAddress(in *models.UpdateAddressModel, id *string) error
	CheckAddressId(id *string) error
	GetAddressAll(id *string) ([]*models.AddressModel, error)
	DeleteAddress(id *string) error
	GetProvince() ([]*entities.Province, error)
	GetDistrictByProvinceId(provinceId *string) ([]*entities.District, error)
	GetSubDistrictByDistrictId(districtId *string) ([]*entities.SubDistrict, error)
}

type addressUsecaseImpl struct {
	addressRepository repositories.AddressRepository
}

func NewAddressUsecaseImpl(addressRepository repositories.AddressRepository) AddressUsecase {
	return &addressUsecaseImpl{
		addressRepository: addressRepository,
	}
}

func (u *addressUsecaseImpl) CheckAddressId(id *string) error {
	if result, err := u.addressRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &addressError.ServerInternalError{Err: err}
		}
		return &addressError.AddressNotFoundError{}
	}
	return nil
}

func (u *addressUsecaseImpl) GetAddressById(id *string) (*models.AddressModel, error) {
	addressData, err := u.addressRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	subDistrictModel := &models.SubDistrict{
		ID:       uint64(addressData.SubDistrict.Id),
		NameTh:   addressData.SubDistrict.NameTh,
		NameEn:   addressData.SubDistrict.NameEn,
		PostCode: addressData.SubDistrict.PostCode,
	}

	districtModel := &models.District{
		ID:     uint64(addressData.District.Id),
		NameTh: addressData.District.NameTh,
		NameEn: addressData.District.NameEn,
	}

	provinceModel := &models.Province{
		ID:     uint64(addressData.Province.Id),
		NameTh: addressData.Province.NameTh,
		NameEn: addressData.Province.NameEn,
	}

	addressModel := &models.AddressModel{
		ID:          uint64(addressData.ID),
		UserId:      uint64(addressData.UserId),
		AddressLine: addressData.AddressLine,
		Phone:       addressData.Phone,
		Name:        addressData.Name,
		SubDistrict: *subDistrictModel,
		District:    *districtModel,
		Province:    *provinceModel,
		Default:     addressData.Default,
		Status:      addressData.Status,
		CreatedAt:   addressData.CreatedAt,
		UpdatedAt:   addressData.UpdatedAt,
	}
	return addressModel, nil
}

func (u *addressUsecaseImpl) GetAddressAll(id *string) ([]*models.AddressModel, error) {
	addresss, err := u.addressRepository.GetDataAllByKey("user_id", id)
	if err != nil {
		return nil, err
	}
	addressModels := []*models.AddressModel{}
	for _, address := range addresss {

		subDistrictModel := &models.SubDistrict{
			ID:       uint64(address.SubDistrict.Id),
			NameTh:   address.SubDistrict.NameTh,
			NameEn:   address.SubDistrict.NameEn,
			PostCode: address.SubDistrict.PostCode,
		}

		districtModel := &models.District{
			ID:     uint64(address.District.Id),
			NameTh: address.District.NameTh,
			NameEn: address.District.NameEn,
		}

		provinceModel := &models.Province{
			ID:     uint64(address.Province.Id),
			NameTh: address.Province.NameTh,
			NameEn: address.Province.NameEn,
		}

		addressModel := &models.AddressModel{
			ID:          uint64(address.ID),
			UserId:      uint64(address.UserId),
			AddressLine: address.AddressLine,
			Phone:       address.Phone,
			Name:        address.Name,
			SubDistrict: *subDistrictModel,
			District:    *districtModel,
			Province:    *provinceModel,
			Default:     address.Default,
			Status:      address.Status,
			CreatedAt:   address.CreatedAt,
			UpdatedAt:   address.UpdatedAt,
		}
		addressModels = append(addressModels, addressModel)
	}
	return addressModels, nil
}

func (u *addressUsecaseImpl) InsertAddress(in *models.InsertAddressModel, id *string) error {
	userId, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	addressInsert := &entities.InsertAddress{
		UserId:        userId,
		AddressLine:   stringIsNil(in.AddressLine),
		Phone:         in.Phone,
		Name:          in.Name,
		SubDistrictId: int(in.SubDistrictId),
		DistrictId:    int(in.DistrictId),
		ProvinceId:    int(in.ProvinceId),
		Default:       in.Default,
		Status:        "active",
	}

	if err := u.addressRepository.InsertData(addressInsert); err != nil {
		return err
	}
	return nil

}

func (u *addressUsecaseImpl) UpdateAddress(in *models.UpdateAddressModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &addressError.ServerInternalError{Err: err}
	}
	addressUpdate := &entities.UpdateAddress{

		AddressLine:   stringIsNil(in.AddressLine),
		Phone:         in.Phone,
		Name:          in.Name,
		SubDistrictId: int(in.SubDistrictId),
		DistrictId:    int(in.DistrictId),
		ProvinceId:    int(in.ProvinceId),
		Default:       in.Default,
		Status:        "active",
	}

	if err := u.addressRepository.UpdateData(addressUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *addressUsecaseImpl) DeleteAddress(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &addressError.ServerInternalError{Err: err}
	}
	if err := u.addressRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}

func (u *addressUsecaseImpl) GetProvince() ([]*entities.Province, error) {
	provinces, err := u.addressRepository.GetProvince()
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (u *addressUsecaseImpl) GetDistrictByProvinceId(provinceId *string) ([]*entities.District, error) {
	districts, err := u.addressRepository.GetDistrictByProvinceId(provinceId)
	if err != nil {
		return nil, err
	}
	return districts, nil
}

func (u *addressUsecaseImpl) GetSubDistrictByDistrictId(districtId *string) ([]*entities.SubDistrict, error) {
	subDistricts, err := u.addressRepository.GetSubDistrictByDistrictId(districtId)
	if err != nil {
		return nil, err
	}
	return subDistricts, nil
}
