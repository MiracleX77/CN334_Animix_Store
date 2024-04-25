package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/delivery/entities"
	deliveryError "github.com/MiracleX77/CN334_Animix_Store/delivery/errors"
	"github.com/MiracleX77/CN334_Animix_Store/delivery/models"
	"github.com/MiracleX77/CN334_Animix_Store/delivery/repositories"
)

type DeliveryUsecase interface {
	InsertDelivery(in *models.InsertDeliveryModel) error
	GetDeliveryById(id *string, token *string) (*models.DeliveryModel, error)
	UpdateDelivery(in *models.UpdateDeliveryModel, id *string) error
	CheckDeliveryId(id *string) error
	GetDeliveryAll(token *string) ([]*models.DeliveryModel, error)
	DeleteDelivery(id *string) error
}

type deliveryUsecaseImpl struct {
	deliveryRepository repositories.DeliveryRepository
}

func NewDeliveryUsecaseImpl(deliveryRepository repositories.DeliveryRepository) DeliveryUsecase {
	return &deliveryUsecaseImpl{
		deliveryRepository: deliveryRepository,
	}
}

func (u *deliveryUsecaseImpl) CheckDeliveryId(id *string) error {
	if result, err := u.deliveryRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &deliveryError.ServerInternalError{Err: err}
		}
		return &deliveryError.DeliveryNotFoundError{}
	}
	return nil
}

func (u *deliveryUsecaseImpl) GetDeliveryById(id *string, token *string) (*models.DeliveryModel, error) {
	deliveryData, err := u.deliveryRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	addressModel := &models.AddressModel{}
	addressId := strconv.Itoa(deliveryData.AddressId)
	if err := getDataFormAPI("5003", "address", addressId, addressModel, *token); err != nil {
		return nil, &deliveryError.ServerInternalError{Err: err}
	}

	deliveryModel := &models.DeliveryModel{
		ID:             uint64(deliveryData.ID),
		Address:        *addressModel,
		Cost:           deliveryData.Cost,
		Type:           deliveryData.Type,
		TrackingNumber: deliveryData.TrackingNumber,
		CreatedAt:      deliveryData.CreatedAt,
		UpdatedAt:      deliveryData.UpdatedAt,
		Status:         deliveryData.Status,
	}

	return deliveryModel, nil
}

func (u *deliveryUsecaseImpl) GetDeliveryAll(token *string) ([]*models.DeliveryModel, error) {
	deliverys, err := u.deliveryRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	deliveryModels := []*models.DeliveryModel{}
	for _, delivery := range deliverys {

		addressModel := &models.AddressModel{}
		addressId := strconv.Itoa(delivery.AddressId)
		if err := getDataFormAPI("5003", "address", addressId, addressModel, *token); err != nil {
			return nil, err
		}

		deliveryModel := &models.DeliveryModel{
			ID:             uint64(delivery.ID),
			Address:        *addressModel,
			Cost:           delivery.Cost,
			Type:           delivery.Type,
			TrackingNumber: delivery.TrackingNumber,
			CreatedAt:      delivery.CreatedAt,
			UpdatedAt:      delivery.UpdatedAt,
			Status:         delivery.Status,
		}
		deliveryModels = append(deliveryModels, deliveryModel)
	}
	return deliveryModels, nil
}

func (u *deliveryUsecaseImpl) InsertDelivery(in *models.InsertDeliveryModel) error {

	deliveryInsert := &entities.InsertDelivery{
		AddressId: int(in.AddressId),
		Status:    "active",
	}

	if _, err := u.deliveryRepository.InsertData(deliveryInsert); err != nil {
		return err
	}
	return nil

}

func (u *deliveryUsecaseImpl) UpdateDelivery(in *models.UpdateDeliveryModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &deliveryError.ServerInternalError{Err: err}
	}
	deliveryUpdate := &entities.UpdateDelivery{
		AddressId:      int(in.AddressId),
		Cost:           in.Cost,
		Type:           in.Type,
		TrackingNumber: in.TrackingNumber,
		Status:         "Processed",
	}
	if err := u.deliveryRepository.UpdateData(deliveryUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *deliveryUsecaseImpl) DeleteDelivery(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &deliveryError.ServerInternalError{Err: err}
	}
	if err := u.deliveryRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
