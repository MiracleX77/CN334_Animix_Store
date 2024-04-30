package server

import (
	reviewHandler "github.com/MiracleX77/CN334_Animix_Store/review/handlers"
	reviewRepository "github.com/MiracleX77/CN334_Animix_Store/review/repositories"
	reviewUsecase "github.com/MiracleX77/CN334_Animix_Store/review/usecases"

	authRepository "github.com/MiracleX77/CN334_Animix_Store/auth/repositories"
)

func (s *echoServer) initializeRouters() {
	s.initializeReviewHttpHandler()
}

func (s *echoServer) initializeReviewHttpHandler() {
	reviewPosgresRepository := reviewRepository.NewReviewPostgresRepository(s.db)
	reviewUsecase := reviewUsecase.NewReviewUsecaseImpl(reviewPosgresRepository)
	reviewHandler := reviewHandler.NewReviewHttpHandler(reviewUsecase)
	reviewRouters := s.app.Group("v1/review")

	reviewRouters.GET("/product/:id", reviewHandler.GetReviewByProductId)
	reviewAuthRouters := s.app.Group("v1/review")

	reviewAuthRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	reviewAuthRouters.GET("/:id", reviewHandler.GetReviewById)
	reviewAuthRouters.GET("/user/:id", reviewHandler.GetReviewByUserId)
	reviewAuthRouters.POST("/", reviewHandler.InsertReview)
	reviewAuthRouters.PUT("/:id", reviewHandler.UpdateReview)
	reviewAuthRouters.DELETE("/:id", reviewHandler.DeleteReview)

}

func authRepositoryForAuth(s *echoServer) authRepository.UserRepository {
	return authRepository.NewUserPostgresRepository(s.db)
}
