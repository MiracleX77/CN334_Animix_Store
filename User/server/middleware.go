package server

import (
	"net/http"
	"strconv"

	tokenUsecase "github.com/MiracleX77/CN334_Animix_Store/auth/usecases"
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	userRepository "github.com/MiracleX77/CN334_Animix_Store/user/repositories"

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
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			} else {
				authHeader = authHeader[7:]
			}

			tokenUsecase := tokenUsecase.NewTokenUsecaseImpl(configs.GetJwtConfig().SecretKey)
			userID, err := tokenUsecase.ParseToken(&authHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			}
			userId := strconv.FormatUint(uint64(*userID), 10)
			result, err := repo.GetUserDataByKey("id", &userId)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			}
			if typeUser == "admin" {
				if result.Type != "Admin" {
					return c.JSON(http.StatusUnauthorized, "invalid or expired token")
				}
			}

			c.Set("userId", userId)
			c.Set("token", authHeader)
			return next(c)
		}
	}
}
