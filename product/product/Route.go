package product

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	product := &Product{}
	e.GET("/products", product.Get)
	e.GET("/products/:id", product.FindById)
	e.POST("/products", product.Persist)
	e.PUT("/products/:id", product.Put)
	e.DELETE("/products/:id", product.Remove)
	return e
}
