package service

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/mkamadeus/grpc-demo/server/models"
)

type weatherService interface {
	GetByLocation(location string) (*models.Weather, error)
}

func New() weatherService {
	return &WeatherService{}
}

type WeatherService struct{}

func (service WeatherService) GetByLocation(location string) (*models.Weather, error) {
	response, err := http.Get("wttr.in/Jakarta?format=\"%%l_%%C_%%t_%%f\"")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	weatherData := strings.Split(string(body), "_")
	if len(weatherData) == 4 {
		return nil, errors.New("invalid weather data")
	}

	return &models.Weather{
		Location:             weatherData[0],
		Condition:            weatherData[1],
		ActualTemperature:    weatherData[2],
		FeelsLikeTemperature: weatherData[3],
	}, nil
}
