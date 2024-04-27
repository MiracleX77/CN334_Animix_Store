package usecases

import (
	"strconv"

	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/repositories"
)

type FavoriteUsecase interface {
	InsertFavorite(in *models.InsertFavoriteModel) error
	GetFavoriteAllByUserId(id *string) ([]*models.FavoriteModel, error)
	DeleteFavorite(id *string) error
}

type favoriteUsecaseImpl struct {
	favoriteRepository repositories.FavoriteRepository
}

func NewFavoriteUsecaseImpl(favoriteRepository repositories.FavoriteRepository) FavoriteUsecase {
	return &favoriteUsecaseImpl{
		favoriteRepository: favoriteRepository,
	}
}

func (u *favoriteUsecaseImpl) GetFavoriteAllByUserId(id *string) ([]*models.FavoriteModel, error) {
	favorites, err := u.favoriteRepository.GetDataAllByKey("user_id", id)
	if err != nil {
		return nil, err
	}
	favoriteModels := []*models.FavoriteModel{}
	for _, favorite := range favorites {
		productModel := &models.ProductFavoriteModel{
			ID:     uint64(favorite.Product.ID),
			Name:   favorite.Product.Name,
			Price:  favorite.Product.Price,
			ImgUrl: favorite.Product.ImgUrl,
		}

		favoriteModel := &models.FavoriteModel{
			ID:        uint64(favorite.ID),
			UserId:    uint64(favorite.UserId),
			Product:   *productModel,
			Status:    favorite.Status,
			CreatedAt: favorite.CreatedAt,
		}

		favoriteModels = append(favoriteModels, favoriteModel)
	}
	return favoriteModels, nil
}

func (u *favoriteUsecaseImpl) InsertFavorite(in *models.InsertFavoriteModel) error {

	authorInsert := &entities.InsertFavorite{
		UserId:    int(in.UserId),
		ProductId: int(in.ProductId),
		Status:    in.Status,
	}

	if err := u.favoriteRepository.InsertData(authorInsert); err != nil {
		return err
	}
	return nil

}

func (u *favoriteUsecaseImpl) DeleteFavorite(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &productError.ServerInternalError{Err: err}
	}
	if err := u.favoriteRepository.DeleteData(&idUint64); err != nil {
		return err
	}
	return nil
}
