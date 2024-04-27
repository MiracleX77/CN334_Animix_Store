package server

import (
	"net/http"
	"strconv"

	userRepository "github.com/MiracleX77/CN334_Animix_Store/auth/repositories"
	tokenUsecase "github.com/MiracleX77/CN334_Animix_Store/auth/usecases"
	"github.com/MiracleX77/CN334_Animix_Store/configs"

	"github.com/labstack/echo/v4"
)

func TokenAuthentication(repo userRepository.UserRepository, typeUser string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "missing authorization header")
			}
			// check if token not start with Bearer
			if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token1")
			} else {
				authHeader = authHeader[7:]
			}

			tokenUsecase := tokenUsecase.NewTokenUsecaseImpl(configs.GetJwtConfig().SecretKey)
			userID, err := tokenUsecase.ParseToken(&authHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token2")
			}
			userId := strconv.FormatUint(uint64(*userID), 10)
			result, err := repo.GetUserDataByKey("id", &userId)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token3")
			}
			if typeUser == "admin" {
				if result.Type != "Admin" {
					return c.JSON(http.StatusUnauthorized, "invalid or expired token4")
				}
			}

			c.Set("userId", userId)
			c.Set("token", authHeader)
			return next(c)
		}
	}
}
