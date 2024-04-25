package usecases

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/repositories"
)

type ProductUsecase interface {
	InsertProduct(in *models.InsertProductModel) error
	GetProductById(id *string) (*models.ProductModel, error)
	UpdateProduct(in *models.UpdateProductModel, id *string) error
	CheckProductId(id *string) error
	GetProductAll() ([]*models.ProductModel, error)
	GetProductAllByKey(key string, id *string) ([]*models.ProductModel, error)
	DeleteProduct(id *string) error
	SendFileToApi(file io.Reader, filename string) error
}

type productUsecaseImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductUsecaseImpl(productRepository repositories.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{
		productRepository: productRepository,
	}
}

func (u *productUsecaseImpl) CheckProductId(id *string) error {
	if result, err := u.productRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &productError.ServerInternalError{Err: err}
		}
		return &productError.ProductNotFoundError{}
	}
	return nil
}

func (u *productUsecaseImpl) GetProductById(id *string) (*models.ProductModel, error) {
	productData, err := u.productRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	authorModel := &models.AuthorModel{
		ID:          uint64(productData.Author.ID),
		Name:        productData.Author.Name,
		Description: productData.Author.Description,
	}
	publisherModel := &models.PublisherModel{
		ID:   uint64(productData.Publisher.ID),
		Name: productData.Publisher.Name,
	}
	categoryModel := &models.CategoryModel{
		ID:   uint64(productData.Category.ID),
		Name: productData.Category.Name,
	}
	productModel := &models.ProductModel{
		ID:          uint64(productData.ID),
		Author:      *authorModel,
		Publisher:   *publisherModel,
		Category:    *categoryModel,
		Name:        productData.Name,
		Description: productData.Description,
		Price:       productData.Price,
		Stock:       productData.Stock,
		ImgUrl:      productData.ImgUrl,
		Status:      productData.Status,
		CreatedAt:   productData.CreatedAt,
		UpdatedAt:   productData.UpdatedAt,
	}

	return productModel, nil
}

func (u *productUsecaseImpl) GetProductAll() ([]*models.ProductModel, error) {
	products, err := u.productRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	productModels := []*models.ProductModel{}
	for _, product := range products {

		authorModel := &models.AuthorModel{
			ID:          uint64(product.Author.ID),
			Name:        product.Author.Name,
			Description: product.Author.Description,
		}
		publisherModel := &models.PublisherModel{
			ID:   uint64(product.Publisher.ID),
			Name: product.Publisher.Name,
		}
		categoryModel := &models.CategoryModel{
			ID:   uint64(product.Category.ID),
			Name: product.Category.Name,
		}
		productModel := &models.ProductModel{
			ID:          uint64(product.ID),
			Author:      *authorModel,
			Publisher:   *publisherModel,
			Category:    *categoryModel,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			ImgUrl:      product.ImgUrl,
			Status:      product.Status,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		productModels = append(productModels, productModel)
	}
	return productModels, nil
}

func (u *productUsecaseImpl) GetProductAllByKey(key string, id *string) ([]*models.ProductModel, error) {
	products, err := u.productRepository.GetDataAllByKey(key, id)
	if err != nil {
		return nil, err
	}
	productModels := []*models.ProductModel{}
	for _, product := range products {

		authorModel := &models.AuthorModel{
			ID:          uint64(product.Author.ID),
			Name:        product.Author.Name,
			Description: product.Author.Description,
		}
		publisherModel := &models.PublisherModel{
			ID:   uint64(product.Publisher.ID),
			Name: product.Publisher.Name,
		}
		categoryModel := &models.CategoryModel{
			ID:   uint64(product.Category.ID),
			Name: product.Category.Name,
		}
		productModel := &models.ProductModel{
			ID:          uint64(product.ID),
			Author:      *authorModel,
			Publisher:   *publisherModel,
			Category:    *categoryModel,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			ImgUrl:      product.ImgUrl,
			Status:      product.Status,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		productModels = append(productModels, productModel)
	}
	return productModels, nil
}

func (u *productUsecaseImpl) InsertProduct(in *models.InsertProductModel) error {

	productInsert := &entities.InsertProduct{
		AuthorId:    int(in.AuthorId),
		CategoryId:  int(in.CategoryId),
		PublisherId: int(in.PublisherId),
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Stock:       in.Stock,
		ImgUrl:      "http://127.0.0.1:8000/images/" + in.Img,
		Status:      "active",
	}

	if err := u.productRepository.InsertData(productInsert); err != nil {
		return err
	}
	return nil

}

func (u *productUsecaseImpl) UpdateProduct(in *models.UpdateProductModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	productUpdate := &entities.UpdateProduct{
		AuthorId:    int(in.AuthorId),
		CategoryId:  int(in.CategoryId),
		PublisherId: int(in.PublisherId),
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Stock:       in.Stock,
		ImgUrl:      "http://127.0.0.1:8000/images/" + in.Img,
		Status:      "active",
	}
	if err := u.productRepository.UpdateData(productUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *productUsecaseImpl) DeleteProduct(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	if err := u.productRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}

func (u *productUsecaseImpl) SendFileToApi(file io.Reader, filename string) error {
	url := "http://127.0.0.1:8000/upload" // Endpoint to which you want to send the file

	// Creating a multipart/form-data body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	writer.Close()

	// Create the request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload file, status code: %d", resp.StatusCode)
	}
	return nil
}
