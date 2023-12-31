package routes

import (
	"server/handlers"
	postgres "server/pkg/database"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRouter(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(postgres.DB)
	h := handlers.HandlerProduct(productRepository)

	e.POST("/product", h.CreateNewProduct)
	e.GET("/products", h.GetAllProduct)
	e.GET("/product/:id", h.GetProductById)
	e.PATCH("/product/:id", middleware.UploadFile(h.UpdateDataProduct))
	e.DELETE("/product/:id", h.DeleteDataProduct)
}
