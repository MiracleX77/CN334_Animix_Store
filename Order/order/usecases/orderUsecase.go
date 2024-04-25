package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/order/entities"
	orderError "github.com/MiracleX77/CN334_Animix_Store/order/errors"
	"github.com/MiracleX77/CN334_Animix_Store/order/models"
	"github.com/MiracleX77/CN334_Animix_Store/order/repositories"

	deliveryRepository "github.com/MiracleX77/CN334_Animix_Store/delivery/repositories"
	paymentRepository "github.com/MiracleX77/CN334_Animix_Store/payment/repositories"

	deliveryModels "github.com/MiracleX77/CN334_Animix_Store/delivery/models"
	paymentModels "github.com/MiracleX77/CN334_Animix_Store/payment/models"

	deliveryEntities "github.com/MiracleX77/CN334_Animix_Store/delivery/entities"
	paymentEntities "github.com/MiracleX77/CN334_Animix_Store/payment/entities"
)

type OrderUsecase interface {
	InsertOrder(in *models.InsertOrderModel) error
	GetOrderById(id *string, token *string) (*models.OrderModel, error)
	UpdateOrder(in *models.UpdateOrderModel, id *string) error
	CheckOrderId(id *string) error
	GetOrderAll(token *string) ([]*models.OrderModel, error)
	DeleteOrder(id *string) error
}

type orderUsecaseImpl struct {
	orderRepository    repositories.OrderRepository
	deliveryRepository deliveryRepository.DeliveryRepository
	paymentRepository  paymentRepository.PaymentRepository
}

func NewOrderUsecaseImpl(orderRepository repositories.OrderRepository, deliveryRepository deliveryRepository.DeliveryRepository, paymentRepository paymentRepository.PaymentRepository) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository:    orderRepository,
		deliveryRepository: deliveryRepository,
		paymentRepository:  paymentRepository,
	}
}

func (u *orderUsecaseImpl) CheckOrderId(id *string) error {
	if result, err := u.orderRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &orderError.ServerInternalError{Err: err}
		}
		return &orderError.OrderNotFoundError{}
	}
	return nil
}

func (u *orderUsecaseImpl) GetOrderById(id *string, token *string) (*models.OrderModel, error) {
	orderData, err := u.orderRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}

	deliveryId := strconv.Itoa(orderData.DeliveryId)
	delivery, err := u.deliveryRepository.GetDataByKey("id", &deliveryId)
	if err != nil {
		return nil, &orderError.ServerInternalError{Err: err}
	}
	addressModel := &deliveryModels.AddressModel{}
	addressId := strconv.Itoa(delivery.AddressId)
	if err := getDataFormAPI("5003", "address", addressId, addressModel, *token); err != nil {
		return nil, &orderError.ServerInternalError{Err: err}
	}
	deliveryModel := &deliveryModels.DeliveryModel{
		ID:             uint64(delivery.ID),
		Address:        *addressModel,
		Cost:           delivery.Cost,
		Type:           delivery.Type,
		TrackingNumber: delivery.TrackingNumber,
		CreatedAt:      delivery.CreatedAt,
		UpdatedAt:      delivery.UpdatedAt,
		Status:         delivery.Status,
	}

	paymentId := strconv.Itoa(orderData.PaymentId)
	payment, err := u.paymentRepository.GetDataByKey("id", &paymentId)
	if err != nil {
		return nil, &orderError.ServerInternalError{Err: err}
	}
	paymentModel := &paymentModels.PaymentModel{
		ID:           uint64(payment.ID),
		Type:         payment.Type,
		Total:        payment.Total,
		ProofPayment: payment.ProofPayment,
		CreatedAt:    payment.CreatedAt,
		UpdatedAt:    payment.UpdatedAt,
		Status:       payment.Status,
	}

	orderModel := &models.OrderModel{
		ID:         uint64(orderData.ID),
		UserId:     uint64(orderData.UserId),
		Delivery:   *deliveryModel,
		Payment:    *paymentModel,
		TotalPrice: orderData.TotalPrice,
		CreatedAt:  orderData.CreatedAt,
		UpdatedAt:  orderData.UpdatedAt,
		Status:     orderData.Status,
	}

	return orderModel, nil
}

func (u *orderUsecaseImpl) GetOrderAll(token *string) ([]*models.ListOrderModel, error) {
	orders, err := u.orderRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	orderModels := []*models.ListOrderModel{}
	for _, order := range orders {
		orderModel := &models.ListOrderModel{
			ID:         uint64(order.ID),
			UserId:     uint64(order.UserId),
			TotalPrice: order.TotalPrice,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
			Status:     order.Status,
		}
		orderModels = append(orderModels, orderModel)
	}
	return orderModels, nil
}

func (u *orderUsecaseImpl) InsertOrder(in *models.InsertOrderModel) error {

	deliveryModels := &deliveryEntities.InsertDelivery{
		AddressId: int(in.AddressId),
		Status:    "Active",
	}

	deliveryId, err := u.deliveryRepository.InsertData(deliveryModels)
	if err != nil {
		return err
	}

	paymentModels := &paymentEntities.InsertPayment{
		Type:         in.Type,
		Total:        in.Total,
		ProofPayment: in.Img,
		Status:       "Active",
	}
	paymentId, err := u.paymentRepository.InsertData(paymentModels)
	if err != nil {
		return err
	}

	orderInsert := &entities.InsertOrder{
		UserId:     int(in.UserId),
		DeliveryId: int(deliveryId),
		PaymentId:  int(paymentId),
		Status:     "Active",
	}

	if err := u.orderRepository.InsertData(orderInsert); err != nil {
		return err
	}
	return nil

}

func (u *orderUsecaseImpl) UpdateOrder(in *models.UpdateOrderModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &orderError.ServerInternalError{Err: err}
	}
	orderUpdate := &entities.UpdateOrder{
		AddressId:      int(in.AddressId),
		Cost:           in.Cost,
		Type:           in.Type,
		TrackingNumber: in.TrackingNumber,
		Status:         "Processed",
	}
	if err := u.orderRepository.UpdateData(orderUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *orderUsecaseImpl) DeleteOrder(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &orderError.ServerInternalError{Err: err}
	}
	if err := u.orderRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
