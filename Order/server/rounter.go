package server

import (
	deliveryHandler "github.com/MiracleX77/CN334_Animix_Store/delivery/handlers"
	deliveryRepository "github.com/MiracleX77/CN334_Animix_Store/delivery/repositories"
	deliveryUsecase "github.com/MiracleX77/CN334_Animix_Store/delivery/usecases"

	paymentHandler "github.com/MiracleX77/CN334_Animix_Store/payment/handlers"
	paymentRepository "github.com/MiracleX77/CN334_Animix_Store/payment/repositories"
	paymentUsecase "github.com/MiracleX77/CN334_Animix_Store/payment/usecases"

	authRepository "github.com/MiracleX77/CN334_Animix_Store/auth/repositories"
)

func (s *echoServer) initializeRouters() {
	s.initializeDeliveryHttpHandler()
	s.initializePaymentHttpHandler()
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

func authRepositoryForAuth(s *echoServer) authRepository.UserRepository {
	return authRepository.NewUserPostgresRepository(s.db)
}
