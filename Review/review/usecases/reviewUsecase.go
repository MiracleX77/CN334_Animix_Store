package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/review/entities"
	reviewError "github.com/MiracleX77/CN334_Animix_Store/review/errors"
	"github.com/MiracleX77/CN334_Animix_Store/review/models"
	"github.com/MiracleX77/CN334_Animix_Store/review/repositories"
)

type ReviewUsecase interface {
	InsertReview(in *models.InsertReviewModel) error
	GetReviewById(id *string, token *string) (*models.ReviewModel, error)
	UpdateReview(in *models.UpdateReviewModel, id *string) error
	CheckReviewId(id *string) error
	GetReviewByKey(key string, value string) ([]*models.ReviewModel, error)
	DeleteReview(id *string) error
}

type reviewUsecaseImpl struct {
	reviewRepository repositories.ReviewRepository
}

func NewReviewUsecaseImpl(reviewRepository repositories.ReviewRepository) ReviewUsecase {
	return &reviewUsecaseImpl{
		reviewRepository: reviewRepository,
	}
}

func (u *reviewUsecaseImpl) CheckReviewId(id *string) error {
	if result, err := u.reviewRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &reviewError.ServerInternalError{Err: err}
		}
		return &reviewError.ReviewNotFoundError{}
	}
	return nil
}

func (u *reviewUsecaseImpl) GetReviewById(id *string, token *string) (*models.ReviewModel, error) {
	reviewData, err := u.reviewRepository.GetDataByKey("id", id)
	if err != nil {
		return nil, err
	}

	reviewModel := &models.ReviewModel{
		ID:        uint64(reviewData.ID),
		UserId:    uint64(reviewData.UserId),
		ProductId: uint64(reviewData.ProductId),
		Title:     reviewData.Title,
		Content:   reviewData.Content,
		Rating:    reviewData.Rating,
		Polarity:  reviewData.Polarity,
		CreatedAt: reviewData.CreatedAt,
		UpdatedAt: reviewData.UpdatedAt,
		Status:    reviewData.Status,
	}

	return reviewModel, nil
}

func (u *reviewUsecaseImpl) GetReviewAllByUserId() ([]*models.ReviewModel, error) {
	reviews, err := u.reviewRepository.GetDataAll()
	if err != nil {
		return nil, err
	}
	reviewModels := []*models.ReviewModel{}
	for _, review := range reviews {
		reviewModel := &models.ReviewModel{
			ID:        uint64(review.ID),
			UserId:    uint64(review.UserId),
			ProductId: uint64(review.ProductId),
			Title:     review.Title,
			Content:   review.Content,
			Rating:    review.Rating,
			Polarity:  review.Polarity,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
			Status:    review.Status,
		}
		reviewModels = append(reviewModels, reviewModel)
	}
	return reviewModels, nil
}
func (u *reviewUsecaseImpl) GetReviewByKey(key string, value string) ([]*models.ReviewModel, error) {
	reviews, err := u.reviewRepository.GetDataAllByKey(key, &value)
	if err != nil {
		return nil, err
	}
	reviewModels := []*models.ReviewModel{}
	for _, review := range reviews {
		reviewModel := &models.ReviewModel{
			ID:        uint64(review.ID),
			UserId:    uint64(review.UserId),
			ProductId: uint64(review.ProductId),
			Title:     review.Title,
			Content:   review.Content,
			Rating:    review.Rating,
			Polarity:  review.Polarity,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
			Status:    review.Status,
		}
		reviewModels = append(reviewModels, reviewModel)
	}
	return reviewModels, nil
}

func (u *reviewUsecaseImpl) InsertReview(in *models.InsertReviewModel) error {

	polarity, err := getDataFormAPI(in.Content)
	if err != nil {
		return err
	}

	reviewInsert := &entities.InsertReview{
		UserId:    int(in.UserId),
		ProductId: int(in.ProductId),
		Title:     in.Title,
		Content:   in.Content,
		Rating:    in.Rating,
		Polarity:  polarity,
		Status:    "Active",
	}

	errr := u.reviewRepository.InsertData(reviewInsert)
	if errr != nil {
		return errr
	}

	return nil

}

func (u *reviewUsecaseImpl) UpdateReview(in *models.UpdateReviewModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &reviewError.ServerInternalError{Err: err}
	}
	polarity, err := getDataFormAPI(in.Content)
	if err != nil {
		return err
	}
	reviewUpdate := &entities.UpdateReview{
		Title:    in.Title,
		Content:  in.Content,
		Rating:   in.Rating,
		Polarity: polarity,
		Status:   in.Status,
	}
	if err := u.reviewRepository.UpdateData(reviewUpdate, &idUint64); err != nil {
		return err
	}
	return nil
}

func (u *reviewUsecaseImpl) DeleteReview(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &reviewError.ServerInternalError{Err: err}
	}
	if err := u.reviewRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
