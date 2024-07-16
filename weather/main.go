package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherStruct struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	fmt.Println("Enter City Name : ")
	// Reading User Input / City
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err1 := scanner.Err(); err1 != nil {
		fmt.Println("Error During Getting Input From User", err1)
		return
	}

	apiKey := "8656581730a2a59b56595237358d823d"
	city := scanner.Text()
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err2 := http.Get(apiUrl)
	if err2 != nil {
		fmt.Println("Error During Get Weather Detail", err2)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	byteData, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("Error During Read Data in Byte", err3)
	}
	var weatherData WeatherStruct
	err4 := json.Unmarshal(byteData, &weatherData)
	if err4 != nil {
		fmt.Println("Error During Unmarshal Data", err4)
		return
	}

	tempKelvin := weatherData.Main.Temp
	tempCelsius := tempKelvin - 273.15

	fmt.Println(tempCelsius, " degree Celsius")
}
