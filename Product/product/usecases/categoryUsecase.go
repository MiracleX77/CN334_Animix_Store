package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/repositories"
)

type CategoryUsecase interface {
	InsertCategory(in *models.InsertCategoryModel) error
	GetCategoryById(id *string) (*models.CategoryModel, error)
	UpdateCategory(in *models.InsertCategoryModel, id *string) error
	GetCategoryAll() ([]*models.CategoryModel, error)
	DeleteCategory(id *string) error
}

type categoryUsecaseImpl struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryUsecaseImpl(categoryRepository repositories.CategoryRepository) CategoryUsecase {
	return &categoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u *categoryUsecaseImpl) GetCategoryById(id *string) (*models.CategoryModel, error) {
	categoryData, err := u.categoryRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	categoryModel := &models.CategoryModel{
		ID:   uint64(categoryData.ID),
		Name: categoryData.Name,
	}

	return categoryModel, nil
}

func (u *categoryUsecaseImpl) GetCategoryAll() ([]*models.CategoryModel, error) {
	categorys, err := u.categoryRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	categoryModels := []*models.CategoryModel{}
	for _, category := range categorys {

		categoryModel := &models.CategoryModel{
			ID:   uint64(category.ID),
			Name: category.Name,
		}

		categoryModels = append(categoryModels, categoryModel)
	}
	return categoryModels, nil
}

func (u *categoryUsecaseImpl) InsertCategory(in *models.InsertCategoryModel) error {

	authorInsert := &entities.InsertCategory{
		Name: in.Name,
	}

	if err := u.categoryRepository.InsertData(authorInsert); err != nil {
		return err
	}
	return nil

}

func (u *categoryUsecaseImpl) UpdateCategory(in *models.InsertCategoryModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	categoryUpdate := &entities.UpdateCategory{
		Name: in.Name,
	}
	if err := u.categoryRepository.UpdateData(categoryUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *categoryUsecaseImpl) DeleteCategory(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	if err := u.categoryRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
