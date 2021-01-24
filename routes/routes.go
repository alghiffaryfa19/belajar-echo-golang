package routes

import (
	"net/http"

	"github.com/alghiffaryfa19/echo-rest/app/controllers"
	"github.com/alghiffaryfa19/echo-rest/app/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New();

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello Fauzan")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai, middleware.IsAuthenticated)
	e.PUT("/pegawai", controllers.UpdatePegawai, middleware.IsAuthenticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middleware.IsAuthenticated)
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/generate-has/:password", controllers.GenerateHashPassword)

	return e
}