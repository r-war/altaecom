package api

import (
	"AltaEcom/api/auth"
	"AltaEcom/api/category"
	"AltaEcom/api/middleware"
	"AltaEcom/api/product"
	"AltaEcom/api/user"
	"AltaEcom/config"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.Controller,
	authController *auth.Controller,
	productController *product.Controller,
	categoryController *category.Controller,
	cfg *config.AppConfig,
) {
	if authController == nil || userController == nil  {
		panic("Controller parameter cannot be nil")
	}

	//authentication with Versioning endpoint
	authV1 := e.Group("api/v1/auth")
	authV1.POST("/login", authController.Login)
	authV1.POST("/register-admin", authController.RegisterAdmin)
	authV1.POST("/register-user", authController.RegisterUser)

	//user with Versioning endpoint
	userV1 := e.Group("api/v1/users")
	userV1.Use(middleware.JWTMiddleware(*cfg))
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("", userController.InsertUser)
	userV1.PUT("/:id", userController.UpdateUser)


	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

}
