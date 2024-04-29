package usecases

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/order/entities"
	orderError "github.com/MiracleX77/CN334_Animix_Store/order/errors"
	"github.com/MiracleX77/CN334_Animix_Store/order/models"
	"github.com/MiracleX77/CN334_Animix_Store/order/repositories"

	deliveryRepository "github.com/MiracleX77/CN334_Animix_Store/delivery/repositories"
	paymentRepository "github.com/MiracleX77/CN334_Animix_Store/payment/repositories"
	transactionRepository "github.com/MiracleX77/CN334_Animix_Store/transaction/repositories"

	deliveryModels "github.com/MiracleX77/CN334_Animix_Store/delivery/models"
	paymentModels "github.com/MiracleX77/CN334_Animix_Store/payment/models"

	deliveryEntities "github.com/MiracleX77/CN334_Animix_Store/delivery/entities"
	paymentEntities "github.com/MiracleX77/CN334_Animix_Store/payment/entities"
	transactionEntities "github.com/MiracleX77/CN334_Animix_Store/transaction/entities"
)

type OrderUsecase interface {
	InsertOrder(in *models.InsertOrderModel) error
	GetOrderById(id *string, token *string) (*models.OrderModel, error)
	UpdateStatusOrder(in *models.UpdateOrderModel, id *string) error
	CheckOrderId(id *string) error
	GetOrderAll() ([]*models.ListOrderModel, error)
	GetOrderByKey(key string, value string) ([]*models.ListOrderModel, error)
	DeleteOrder(id *string) error
	SendFileToApi(file io.Reader, filename string) error
}

type orderUsecaseImpl struct {
	orderRepository       repositories.OrderRepository
	deliveryRepository    deliveryRepository.DeliveryRepository
	paymentRepository     paymentRepository.PaymentRepository
	transactionRepository transactionRepository.TransactionRepository
}

func NewOrderUsecaseImpl(orderRepository repositories.OrderRepository, deliveryRepository deliveryRepository.DeliveryRepository, paymentRepository paymentRepository.PaymentRepository, transactionRepository transactionRepository.TransactionRepository) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository:       orderRepository,
		deliveryRepository:    deliveryRepository,
		paymentRepository:     paymentRepository,
		transactionRepository: transactionRepository,
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

func (u *orderUsecaseImpl) GetOrderAll() ([]*models.ListOrderModel, error) {
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
func (u *orderUsecaseImpl) GetOrderByKey(key string, value string) ([]*models.ListOrderModel, error) {
	orders, err := u.orderRepository.GetDataAllByKey(key, &value)
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
		ProofPayment: "http://127.0.0.1:8001/images/" + in.Img,
		Status:       "WaitConfirm",
	}
	paymentId, err := u.paymentRepository.InsertData(paymentModels)
	if err != nil {
		return err
	}

	orderInsert := &entities.InsertOrder{
		UserId:     int(in.UserId),
		DeliveryId: int(deliveryId),
		PaymentId:  int(paymentId),
		TotalPrice: in.Total,
		Status:     "Pending",
	}

	orderId, err := u.orderRepository.InsertData(orderInsert)
	if err != nil {
		return err
	}

	for _, productId := range in.ListProductId {
		transactionModels := &transactionEntities.InsertTransaction{
			ProductId: int(productId),
			OrderId:   int(orderId),
			Status:    "Active",
		}
		if _, err := u.transactionRepository.InsertData(transactionModels); err != nil {
			return err
		}
	}

	return nil

}

func (u *orderUsecaseImpl) UpdateStatusOrder(in *models.UpdateOrderModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &orderError.ServerInternalError{Err: err}
	}
	orderUpdate := &entities.UpdateOrder{
		Status: in.Status,
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

func (u *orderUsecaseImpl) SendFileToApi(file io.Reader, filename string) error {
	url := "http://storage:8000/upload" // Endpoint to which you want to send the file

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
