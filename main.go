package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Rectangle struct {
	Width  int `json: "width"`
	Height int `json: "height"`
	Area   int `json: "area"`
}

func (r *Rectangle) rectangle_area() int {
	r.Area = r.Width * r.Height
	return r.Area
}

func calculate(c echo.Context) error {
	var newArea Rectangle
	a := &newArea

	a.Width, _ = strconv.Atoi(c.QueryParam("width"))
	a.Height, _ = strconv.Atoi(c.QueryParam("height"))

	a.rectangle_area()
	return c.JSON(http.StatusOK, a)
}

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", calculate)

	app.Logger.Fatal(app.Start(":8000"))
}
