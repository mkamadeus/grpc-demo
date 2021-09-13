package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mkamadeus/grpc-demo/server/models"
	"github.com/mkamadeus/grpc-demo/server/schemas"
)

func NewWeatherController() models.WeatherController {
	return &WeatherController{}
}

type WeatherController struct{}

func (*WeatherController) GetByLocation(location string) (*schemas.WeatherReply, error) {
	response, err := http.Get("wttr.in/Jakarta?format=\"%%l_%%C_%%t_%%f\"")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	weatherData := strings.Split(string(body), "_")
	if len(weatherData) == 4 {
		return nil, errors.New("invalid weather data")
	}

	return &schemas.WeatherReply{
		Location:             weatherData[0],
		Condition:            weatherData[1],
		ActualTemperature:    weatherData[2],
		FeelsLikeTemperature: weatherData[3],
	}, nil
}
