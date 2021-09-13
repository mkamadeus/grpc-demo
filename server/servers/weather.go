package servers

import (
	"context"

	"github.com/mkamadeus/grpc-demo/server/models"
	"github.com/mkamadeus/grpc-demo/server/schemas"
)

type WeatherServer struct {
	WeatherCtl models.WeatherController
	schemas.UnimplementedWeatherServer
}

func (server *WeatherServer) GetWeatherInformation(ctx context.Context, req *schemas.WeatherRequest) (*schemas.WeatherReply, error) {
	return server.WeatherCtl.GetByLocation(req.GetLocation())
}
