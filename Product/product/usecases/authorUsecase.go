package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/repositories"
)

type AuthorUsecase interface {
	InsertAuthor(in *models.InsertAuthorModel) error
	GetAuthorById(id *string) (*models.AuthorModel, error)
	UpdateAuthor(in *models.InsertAuthorModel, id *string) error
	GetAuthorAll() ([]*models.AuthorModel, error)
	DeleteAuthor(id *string) error
}

type authorUsecaseImpl struct {
	authorRepository repositories.AuthorRepository
}

func NewAuthorUsecaseImpl(authorRepository repositories.AuthorRepository) AuthorUsecase {
	return &authorUsecaseImpl{
		authorRepository: authorRepository,
	}
}

func (u *authorUsecaseImpl) GetAuthorById(id *string) (*models.AuthorModel, error) {
	authData, err := u.authorRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	authorModel := &models.AuthorModel{
		ID:          uint64(authData.ID),
		Name:        authData.Name,
		Description: authData.Description,
	}

	return authorModel, nil
}

func (u *authorUsecaseImpl) GetAuthorAll() ([]*models.AuthorModel, error) {
	authors, err := u.authorRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	authorModels := []*models.AuthorModel{}
	for _, author := range authors {

		authorModel := &models.AuthorModel{
			ID:          uint64(author.ID),
			Name:        author.Name,
			Description: author.Description,
		}

		authorModels = append(authorModels, authorModel)
	}
	return authorModels, nil
}

func (u *authorUsecaseImpl) InsertAuthor(in *models.InsertAuthorModel) error {

	authorInsert := &entities.InsertAuthor{
		Name:        in.Name,
		Description: in.Description,
	}

	if err := u.authorRepository.InsertData(authorInsert); err != nil {
		return err
	}
	return nil

}

func (u *authorUsecaseImpl) UpdateAuthor(in *models.InsertAuthorModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	productUpdate := &entities.UpdateAuthor{
		Name:        in.Name,
		Description: in.Description,
	}
	if err := u.authorRepository.UpdateData(productUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *authorUsecaseImpl) DeleteAuthor(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	if err := u.authorRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
