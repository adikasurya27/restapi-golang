package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Users struct {
	Email  string 
	Nama   string
	Umur   int
	Alamat string
}

func main() {
	route := echo.New()

	route.POST("user/create_user", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)

		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses membuat user baru",
			Data:    *user,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:email", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)
		user.Email = c.Param("email")
		// update ke database
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses mengupdate data user",
			Data:    *user,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_use/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		// delete data
		response := struct {
			Message string
			Data    string
		}{
			Message: "Sukses menghapus data user dengan email",
			Data: 	user.Email,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/get_data", func(c echo.Context) error {
		user := new(Users)
		user.Email = "fbidoira@gmail.com"
		user.Nama = "Adika Surya Perdana"
		user.Umur = 22
		user.Alamat = "Trosobo"

		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses melihat data user",
			Data:    *user,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.Start(":8080")
}
