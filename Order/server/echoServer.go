package server

import (
	"fmt"

	"github.com/MiracleX77/CN334_Animix_Store/configs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *configs.Config
}

func NewEchoServer(cfg *configs.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"*"},
	}))
	// initialize routers here
	s.initializeRouters()
	s.app.Validator = NewCustomValidator()
	s.app.Use(middleware.Logger())
	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func genderValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == "female" || value == "male"
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("gender", genderValidation)
	return &CustomValidator{validator: v}
}
