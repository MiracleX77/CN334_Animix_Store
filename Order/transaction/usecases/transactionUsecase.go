package usecases

import (
	"strconv"

	transactionError "github.com/MiracleX77/CN334_Animix_Store/transaction/errors"
	"github.com/MiracleX77/CN334_Animix_Store/transaction/models"
	"github.com/MiracleX77/CN334_Animix_Store/transaction/repositories"
)

type TransactionUsecase interface {
	GetTransactionById(id *string, token *string) (*models.TransactionModel, error)
	GetTransactionAllByOrderId(id *string, token *string) ([]*models.TransactionModel, error)
}

type transactionUsecaseImpl struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionUsecaseImpl(transactionRepository repositories.TransactionRepository) TransactionUsecase {
	return &transactionUsecaseImpl{
		transactionRepository: transactionRepository,
	}
}

func (u *transactionUsecaseImpl) GetTransactionById(id *string, token *string) (*models.TransactionModel, error) {
	transactionData, err := u.transactionRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	productModel := &models.ProductModel{}
	productId := strconv.Itoa(transactionData.ProductId)
	if err := getDataFormAPI("5002", "product", productId, productModel, *token); err != nil {
		return nil, &transactionError.ServerInternalError{Err: err}
	}

	transactionModel := &models.TransactionModel{
		ID:        uint64(transactionData.ID),
		Product:   *productModel,
		OrderId:   uint64(transactionData.OrderId),
		CreatedAt: transactionData.CreatedAt,
		UpdatedAt: transactionData.UpdatedAt,
		Status:    transactionData.Status,
	}

	return transactionModel, nil
}

func (u *transactionUsecaseImpl) GetTransactionAllByOrderId(id *string, token *string) ([]*models.TransactionModel, error) {
	transactions, err := u.transactionRepository.GetDataAllByKey("order_id", id)
	if err != nil {
		return nil, err
	}
	transactionModels := []*models.TransactionModel{}
	for _, transaction := range transactions {

		productModel := &models.ProductModel{}
		productId := strconv.Itoa(transaction.ProductId)
		if err := getDataFormAPI("5002", "product", productId, productModel, *token); err != nil {
			return nil, &transactionError.ServerInternalError{Err: err}
		}

		transactionModel := &models.TransactionModel{
			ID:        uint64(transaction.ID),
			Product:   *productModel,
			OrderId:   uint64(transaction.OrderId),
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
			Status:    transaction.Status,
		}
		transactionModels = append(transactionModels, transactionModel)
	}
	return transactionModels, nil
}
