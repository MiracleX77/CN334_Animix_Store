package usecases

import (
	"github.com/MiracleX77/CN334_Animix_Store/payment/entities"
	paymentError "github.com/MiracleX77/CN334_Animix_Store/payment/errors"
	"github.com/MiracleX77/CN334_Animix_Store/payment/models"
	"github.com/MiracleX77/CN334_Animix_Store/payment/repositories"
)

type PaymentUsecase interface {
	InsertPayment(in *models.InsertPaymentModel) error
	GetPaymentById(id *string) (*models.PaymentModel, error)
	CheckPaymentId(id *string) error
}

type paymentUsecaseImpl struct {
	paymentRepository repositories.PaymentRepository
}

func NewPaymentUsecaseImpl(paymentRepository repositories.PaymentRepository) PaymentUsecase {
	return &paymentUsecaseImpl{
		paymentRepository: paymentRepository,
	}
}

func (u *paymentUsecaseImpl) CheckPaymentId(id *string) error {
	if result, err := u.paymentRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &paymentError.ServerInternalError{Err: err}
		}
		return &paymentError.PaymentNotFoundError{}
	}
	return nil
}

func (u *paymentUsecaseImpl) GetPaymentById(id *string) (*models.PaymentModel, error) {
	paymentData, err := u.paymentRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}

	paymentModel := &models.PaymentModel{
		ID:           uint64(paymentData.ID),
		Type:         paymentData.Type,
		Total:        paymentData.Total,
		ProofPayment: paymentData.ProofPayment,
		Status:       paymentData.Status,
		CreatedAt:    paymentData.CreatedAt,
	}

	return paymentModel, nil
}

func (u *paymentUsecaseImpl) InsertPayment(in *models.InsertPaymentModel) error {

	paymentInsert := &entities.InsertPayment{
		Type:         in.Type,
		Total:        in.Total,
		ProofPayment: in.ProofPayment,
		Status:       "active",
	}

	if _, err := u.paymentRepository.InsertData(paymentInsert); err != nil {
		return err
	}
	return nil

}
