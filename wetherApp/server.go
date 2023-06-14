package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
	}))

	e.POST("/weather", getWeatherDetails)

	e.Logger.Fatal(e.Start(":8000"))
}

type City struct {
	City_name string `json:"cityName"`
}

func getWeatherDetails(c echo.Context) error {

	var requestData City

	// Bind the JSON data from the request body into the requestData variable
	if err := c.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON data")
	}

	// Access the values from the requestData variable
	city_name := requestData.City_name
	fmt.Println("city_name: ", city_name)
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
	type WeatherData struct {
		Main struct {
			Temperature float64 `json:"temp"`     //struct field for the temp
			Humidity    int     `json:"humidity"` //struct field for the humidity
			Pressure    int     `json:"pressure"` //struct field for the pressure
		} `json:"main"` //struct the field for the main weatherData with json tag
		Name string `json:"name"` //field for the city name
	}

	var weatherData WeatherData                                //created the variable to get the unmarshaled weather data
	if err := json.Unmarshal(body, &weatherData); err != nil { //unmarshal the json into the weatherdata variable
		return err
	}

	temperatureCelcius := math.Round(weatherData.Main.Temperature - 273.15) // convert the kelvin to celcius

	// Build a response map with the required fields
	responseData := map[string]interface{}{
		"City":        weatherData.Name,
		"Humidity":    weatherData.Main.Humidity,
		"Pressure":    weatherData.Main.Pressure,
		"Temperature": temperatureCelcius,
	}

	return c.JSON(http.StatusOK, responseData)

}
