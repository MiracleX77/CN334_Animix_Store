package server

import (
	deliveryHandler "github.com/MiracleX77/CN334_Animix_Store/delivery/handlers"
	deliveryRepository "github.com/MiracleX77/CN334_Animix_Store/delivery/repositories"
	deliveryUsecase "github.com/MiracleX77/CN334_Animix_Store/delivery/usecases"

	paymentHandler "github.com/MiracleX77/CN334_Animix_Store/payment/handlers"
	paymentRepository "github.com/MiracleX77/CN334_Animix_Store/payment/repositories"
	paymentUsecase "github.com/MiracleX77/CN334_Animix_Store/payment/usecases"

	transactionHandler "github.com/MiracleX77/CN334_Animix_Store/transaction/handlers"
	transactionRepository "github.com/MiracleX77/CN334_Animix_Store/transaction/repositories"
	transactionUsecase "github.com/MiracleX77/CN334_Animix_Store/transaction/usecases"

	orderHandler "github.com/MiracleX77/CN334_Animix_Store/order/handlers"
	orderRepository "github.com/MiracleX77/CN334_Animix_Store/order/repositories"
	orderUsecase "github.com/MiracleX77/CN334_Animix_Store/order/usecases"

	authRepository "github.com/MiracleX77/CN334_Animix_Store/auth/repositories"
)

func (s *echoServer) initializeRouters() {
	s.initializeDeliveryHttpHandler()
	s.initializePaymentHttpHandler()
	s.initializeTransactionHttpHandler()
	s.initializeOrderHttpHandler()
}

func (s *echoServer) initializeDeliveryHttpHandler() {
	deliveryPosgresRepository := deliveryRepository.NewDeliveryPostgresRepository(s.db)
	deliveryUsecase := deliveryUsecase.NewDeliveryUsecaseImpl(deliveryPosgresRepository)
	deliveryHandler := deliveryHandler.NewDeliveryHttpHandler(deliveryUsecase)

	deliveryRouters := s.app.Group("v1/delivery")

	deliveryRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	deliveryRouters.GET("/:id", deliveryHandler.GetDeliveryById)

	adminRouters := s.app.Group("v1/delivery")
	adminRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "admin"))
	adminRouters.GET("/", deliveryHandler.GetDeliveryAll)
	adminRouters.PUT("/:id", deliveryHandler.UpdateDelivery)
	adminRouters.DELETE("/:id", deliveryHandler.DeleteDelivery)

}

func (s *echoServer) initializePaymentHttpHandler() {
	paymentPosgresRepository := paymentRepository.NewPaymentPostgresRepository(s.db)
	paymentUsecase := paymentUsecase.NewPaymentUsecaseImpl(paymentPosgresRepository)
	paymentHandler := paymentHandler.NewPaymentHttpHandler(paymentUsecase)

	paymentRouters := s.app.Group("v1/payment")

	paymentRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	paymentRouters.GET("/:id", paymentHandler.GetPaymentById)

}

func (s *echoServer) initializeTransactionHttpHandler() {
	transactionPosgresRepository := transactionRepository.NewTransactionPostgresRepository(s.db)
	transactionUsecase := transactionUsecase.NewTransactionUsecaseImpl(transactionPosgresRepository)
	transactionHandler := transactionHandler.NewTransactionHttpHandler(transactionUsecase)

	transactionRouters := s.app.Group("v1/transaction")

	transactionRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	transactionRouters.GET("/:id", transactionHandler.GetTransactionById)
	transactionRouters.GET("/order/:id", transactionHandler.GetTransactionAllByOrderId)
}

func (s *echoServer) initializeOrderHttpHandler() {
	orderPosgresRepository := orderRepository.NewOrderPostgresRepository(s.db)
	transactionPosgresRepository := transactionRepository.NewTransactionPostgresRepository(s.db)
	deliveryPosgresRepository := deliveryRepository.NewDeliveryPostgresRepository(s.db)
	paymentPosgresRepository := paymentRepository.NewPaymentPostgresRepository(s.db)
	orderUsecase := orderUsecase.NewOrderUsecaseImpl(orderPosgresRepository, deliveryPosgresRepository, paymentPosgresRepository, transactionPosgresRepository)
	orderHandler := orderHandler.NewOrderHttpHandler(orderUsecase)

	orderRouters := s.app.Group("v1/order")

	orderRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	orderRouters.GET("/:id", orderHandler.GetOrderById)
	orderRouters.GET("/user/:id", orderHandler.GetOrderByUserId)
	orderRouters.POST("/", orderHandler.InsertOrder)

	adminRouters := s.app.Group("v1/order")
	adminRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "admin"))
	adminRouters.GET("/", orderHandler.GetOrderAll)
	adminRouters.GET("/status/:status", orderHandler.GetOrderByStatus)
	adminRouters.PUT("/:id", orderHandler.UpdateOrder)
	adminRouters.DELETE("/:id", orderHandler.DeleteOrder)

}

func authRepositoryForAuth(s *echoServer) authRepository.UserRepository {
	return authRepository.NewUserPostgresRepository(s.db)
}
