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

	reviewRouters.Use(TokenAuthentication(authRepositoryForAuth(s), "user"))
	reviewRouters.GET("/:id", reviewHandler.GetReviewById)
	reviewRouters.GET("/user/:id", reviewHandler.GetReviewByUserId)
	reviewRouters.GET("/product/:id", reviewHandler.GetReviewByProductId)
	reviewRouters.POST("/", reviewHandler.InsertReview)
	reviewRouters.PUT("/:id", reviewHandler.UpdateReview)
	reviewRouters.DELETE("/:id", reviewHandler.DeleteReview)

}

func authRepositoryForAuth(s *echoServer) authRepository.UserRepository {
	return authRepository.NewUserPostgresRepository(s.db)
}
