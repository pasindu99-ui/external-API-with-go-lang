package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	e.GET("/weather/", getWeatherDetails)

	e.Logger.Fatal(e.Start(":8000"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getWeatherDetails(c echo.Context) error {
	city_name := c.FormValue("location")
	API_key := "db8e2e85e965d896159bf85f3880e393"
	API_URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city_name, API_key)

	response, err := http.Get(API_URL)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, body)

}
