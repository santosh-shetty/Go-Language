package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherRes struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		city := c.Query("city")
		if city == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
			return
		}

		// apiKey := ""
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
		res, err := http.Get(url)

		if err != nil {
			log.Println("Error During Weather API call:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
			return
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			var weatherData WeatherRes
			if err := json.Unmarshal(bodyBytes, &weatherData); err != nil {
				log.Fatal(err)
			}

			tempKelvin := weatherData.Main.Temp
			tempCelsius := tempKelvin - 273.15
			tempFahrenheit := (tempKelvin-273.15)*9/5 + 32

			c.JSON(http.StatusOK, gin.H{
				"temperature_kelvin":     tempKelvin,
				"temperature_celsius":    tempCelsius,
				"temperature_fahrenheit": tempFahrenheit,
				"city":                   city,
			})
		} else {
			c.JSON(res.StatusCode, gin.H{"error": "Failed to retrieve weather data"})
		}
	})

	r.Run(":8080") // Start the server on port 8080
}
