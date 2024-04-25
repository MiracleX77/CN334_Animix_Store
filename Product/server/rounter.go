package server

import (
	productHandler "github.com/MiracleX77/CN334_Animix_Store/product/handlers"
	productRepository "github.com/MiracleX77/CN334_Animix_Store/product/repositories"
	productUsecase "github.com/MiracleX77/CN334_Animix_Store/product/usecases"

	authRepository "github.com/MiracleX77/CN334_Animix_Store/auth/repositories"
)

func (s *echoServer) initializeRouters() {
	s.initializeProductHttpHandler()
	s.initializeAuthorHttpHandler()
	s.initializePublisherHttpHandler()
	s.initializeCategoryHttpHandler()

}

func (s *echoServer) initializeProductHttpHandler() {
	productPosgresRepository := productRepository.NewProductPostgresRepository(s.db)
	productUsecase := productUsecase.NewProductUsecaseImpl(productPosgresRepository)
	productHandler := productHandler.NewProductHttpHandler(productUsecase)

	productRouters := s.app.Group("v1/product")
	productRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	productRouters.GET("/:id", productHandler.GetProductById)
	productRouters.GET("/", productHandler.GetProductAll)
	productRouters.GET("/category/:id", productHandler.GetProductAllByCategory)
	productRouters.GET("/search/:name", productHandler.GetProductAllByName)

	adminRouters := s.app.Group("v1/product")
	adminRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "admin"))
	adminRouters.POST("/", productHandler.InsertProduct)
	adminRouters.PUT("/", productHandler.UpdateProduct)
	adminRouters.DELETE("/:id", productHandler.DeleteProduct)

}

func (s *echoServer) initializeAuthorHttpHandler() {
	authorPosgresRepository := productRepository.NewAuthorPostgresRepository(s.db)
	authorUsecase := productUsecase.NewAuthorUsecaseImpl(authorPosgresRepository)
	authorHandler := productHandler.NewAuthorHttpHandler(authorUsecase)

	authorRouters := s.app.Group("v1/author")

	authorRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	authorRouters.GET("/:id", authorHandler.GetAuthorById)
	authorRouters.GET("/", authorHandler.GetAuthorAll)

	adminRouters := s.app.Group("v1/author")
	adminRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "admin"))
	adminRouters.POST("/", authorHandler.InsertAuthor)
	adminRouters.PUT("/:id", authorHandler.UpdateAuthor)
	adminRouters.DELETE("/:id", authorHandler.DeleteAuthor)

}

func (s *echoServer) initializePublisherHttpHandler() {
	publisherPosgresRepository := productRepository.NewPublisherPostgresRepository(s.db)
	publisherUsecase := productUsecase.NewPublisherUsecaseImpl(publisherPosgresRepository)
	publisherHandler := productHandler.NewPublisherHttpHandler(publisherUsecase)

	publisherRouters := s.app.Group("v1/publisher")

	publisherRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	publisherRouters.GET("/:id", publisherHandler.GetPublisherById)
	publisherRouters.GET("/", publisherHandler.GetPublisherAll)

	adminRouters := s.app.Group("v1/publisher")
	adminRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "admin"))
	adminRouters.POST("/", publisherHandler.InsertPublisher)
	adminRouters.PUT("/:id", publisherHandler.UpdatePublisher)
	adminRouters.DELETE("/:id", publisherHandler.DeletePublisher)
}

func (s *echoServer) initializeCategoryHttpHandler() {
	categoryPosgresRepository := productRepository.NewCategoryPostgresRepository(s.db)
	categoryUsecase := productUsecase.NewCategoryUsecaseImpl(categoryPosgresRepository)
	categoryHandler := productHandler.NewCategoryHttpHandler(categoryUsecase)

	categoryRouters := s.app.Group("v1/category")

	categoryRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	categoryRouters.GET("/:id", categoryHandler.GetCategoryById)
	categoryRouters.GET("/", categoryHandler.GetCategoryAll)

	adminRouters := s.app.Group("v1/category")
	adminRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "admin"))
	adminRouters.POST("/", categoryHandler.InsertCategory)
	adminRouters.PUT("/:id", categoryHandler.UpdateCategory)
	adminRouters.DELETE("/:id", categoryHandler.DeleteCategory)

}

func authRepositoryForAuth(s *echoServer) authRepository.UserRepository {
	return authRepository.NewUserPostgresRepository(s.db)
}
