package usecases

import (
	authError "github.com/MiracleX77/CN334_Animix_Store/auth/errors"

	"github.com/dgrijalva/jwt-go"
)

type TokenUsecase interface {
	ParseToken(token *string) (*uint, error)
}

type tokenUsecaseImpl struct {
	secretKey string
}

func NewTokenUsecaseImpl(secretKey string) TokenUsecase {
	return &tokenUsecaseImpl{
		secretKey: secretKey,
	}
}

func (u *tokenUsecaseImpl) ParseToken(token *string) (*uint, error) {
	key := []byte(u.secretKey)
	// Parse the token
	t, err := jwt.Parse(*token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &authError.TokenNotAuthorizedError{}
		}
		return key, nil
	})
	if err != nil {
		return nil, &authError.ServerInternalError{Err: err}
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, &authError.TokenNotAuthorizedError{}
	}
	user_id := uint(claims["user_id"].(float64))
	return &user_id, nil
}
