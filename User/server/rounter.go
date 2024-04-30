package server

import (
	authHandler "github.com/MiracleX77/CN334_Animix_Store/auth/handlers"
	authUsecase "github.com/MiracleX77/CN334_Animix_Store/auth/usecases"

	userHandler "github.com/MiracleX77/CN334_Animix_Store/user/handlers"
	userRepository "github.com/MiracleX77/CN334_Animix_Store/user/repositories"
	userUsecase "github.com/MiracleX77/CN334_Animix_Store/user/usecases"

	addressHandler "github.com/MiracleX77/CN334_Animix_Store/address/handlers"
	addressRepository "github.com/MiracleX77/CN334_Animix_Store/address/repositories"
	addressUsecase "github.com/MiracleX77/CN334_Animix_Store/address/usecases"
)

func (s *echoServer) initializeRouters() {
	s.initializeAuthHttpHandler()
	s.initializeUserHttpHandler()
	s.initializeAddressHttpHandler()
}

func (s *echoServer) initializeAuthHttpHandler() {
	userPosgresRepository := userRepository.NewUserPostgresRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecaseImpl(userPosgresRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(authUsecase)

	authRouters := s.app.Group("v1/auth")

	authRouters.POST("/register", authHttpHandler.Register)
	authRouters.POST("/login", authHttpHandler.Login)
}

func (s *echoServer) initializeUserHttpHandler() {
	userPosgresRepository := userRepository.NewUserPostgresRepository(s.db)
	userUsecase := userUsecase.NewUserUsecaseImpl(userPosgresRepository)
	userHttpHandler := userHandler.NewUserHttpHandler(userUsecase)

	uRouters := s.app.Group("v1/user")
	uRouters.GET("/:id", userHttpHandler.GetUserByUserId)

	userRouters := s.app.Group("v1/user")

	userRouters.Use(TokenAuthentication(userRepositoryForAuth(s), "user"))
	userRouters.PUT("/", userHttpHandler.UpdateUser)
	userRouters.GET("/", userHttpHandler.GetUserById)
	userRouters.DELETE("/", userHttpHandler.DeleteUser)

	adminRouters := s.app.Group("v1/user")
	adminRouters.Use(TokenAuthentication(userRepositoryForAuth(s), "admin"))
	adminRouters.GET("/all", userHttpHandler.GetUserAll)
}

func (s *echoServer) initializeAddressHttpHandler() {
	addressPosgresRepository := addressRepository.NewAddressPostgresRepository(s.db)
	addressUsecase := addressUsecase.NewAddressUsecaseImpl(addressPosgresRepository)
	addressHttpHandler := addressHandler.NewAddressHttpHandler(addressUsecase)

	addressRouters := s.app.Group("v1/address")

	addressRouters.Use(TokenAuthentication(userRepositoryForAuth(s), "user"))
	addressRouters.POST("/", addressHttpHandler.InsertAddress)
	addressRouters.PUT("/", addressHttpHandler.UpdateAddress)
	addressRouters.GET("/:id", addressHttpHandler.GetAddressById)
	addressRouters.GET("/", addressHttpHandler.GetAddressAll)
	addressRouters.DELETE("/:id", addressHttpHandler.DeleteAddress)
	addressRouters.GET("/province", addressHttpHandler.GetProvince)
	addressRouters.GET("/district/:id", addressHttpHandler.GetDistrictByProvinceId)
	addressRouters.GET("/subdistrict/:id", addressHttpHandler.GetSubDistrictByDistrictId)
}

func userRepositoryForAuth(s *echoServer) userRepository.UserRepository {
	return userRepository.NewUserPostgresRepository(s.db)
}
