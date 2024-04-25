package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/repositories"
)

type PublisherUsecase interface {
	InsertPublisher(in *models.InsertPublisherModel) error
	GetPublisherById(id *string) (*models.PublisherModel, error)
	UpdatePublisher(in *models.InsertPublisherModel, id *string) error
	GetPublisherAll() ([]*models.PublisherModel, error)
	DeletePublisher(id *string) error
}

type publisherUsecaseImpl struct {
	publisherRepository repositories.PublisherRepository
}

func NewPublisherUsecaseImpl(publisherRepository repositories.PublisherRepository) PublisherUsecase {
	return &publisherUsecaseImpl{
		publisherRepository: publisherRepository,
	}
}

func (u *publisherUsecaseImpl) GetPublisherById(id *string) (*models.PublisherModel, error) {
	publisherData, err := u.publisherRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	publisherModel := &models.PublisherModel{
		ID:   uint64(publisherData.ID),
		Name: publisherData.Name,
	}

	return publisherModel, nil
}

func (u *publisherUsecaseImpl) GetPublisherAll() ([]*models.PublisherModel, error) {
	publishers, err := u.publisherRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	publisherModels := []*models.PublisherModel{}
	for _, publisher := range publishers {

		publisherModel := &models.PublisherModel{
			ID:   uint64(publisher.ID),
			Name: publisher.Name,
		}

		publisherModels = append(publisherModels, publisherModel)
	}
	return publisherModels, nil
}

func (u *publisherUsecaseImpl) InsertPublisher(in *models.InsertPublisherModel) error {

	authorInsert := &entities.InsertPublisher{
		Name: in.Name,
	}

	if err := u.publisherRepository.InsertData(authorInsert); err != nil {
		return err
	}
	return nil

}

func (u *publisherUsecaseImpl) UpdatePublisher(in *models.InsertPublisherModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	publisherUpdate := &entities.UpdatePublisher{
		Name: in.Name,
	}
	if err := u.publisherRepository.UpdateData(publisherUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *publisherUsecaseImpl) DeletePublisher(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	if err := u.publisherRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
