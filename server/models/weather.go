package models

import (
	"github.com/mkamadeus/grpc-demo/server/schemas"
)

type WeatherController interface {
	GetByLocation(location string) (*schemas.WeatherReply, error)
}
