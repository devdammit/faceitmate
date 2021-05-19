package faceitmate

import (
	"context"
	"errors"
	"github.com/devdammit/faceitmate/pkg/api"
	"github.com/devdammit/faceitmate/pkg/faceit"
	"log"
)

// GRPCServer grpc server
type GRPCServer struct {
	faceitClient faceit.API
}

func NewGRPCServer() *GRPCServer {
	fcOptions := faceit.APIOptions{ApiKey: "f578dd71-6ce7-448a-bf3f-2d265cb0f145"}
	fc := faceit.NewFaceitClient(fcOptions)

	return &GRPCServer{faceitClient: *fc}
}

func (server *GRPCServer) AddPlayer(_ context.Context, req *api.RegisterPlayerRequest) (*api.RegisterWatchingResponse, error) {
	playerParams := faceit.PlayerQuery{
		Nickname: req.Nickname,
	}

	player, err := server.faceitClient.GetPlayer(playerParams)
	if err != nil {
		var re *faceit.ResponseError
		ok := errors.Is(err, re)

		if ok {
			var code api.RegisterWatchingResponse_ResponseCode

			switch {
			case re.NotFound():
				code = api.RegisterWatchingResponse_NOT_FOUND
			case re.NotAuthorized():
				code = api.RegisterWatchingResponse_FAILURE
			case re.Temporary():
				code = api.RegisterWatchingResponse_FACEIT_NOT_RESPONSE
			}

			return &api.RegisterWatchingResponse{
				ResponseCode: code,
				Message:      re.Error(),
				PlayerId:     "",
			}, nil
		}

		return &api.RegisterWatchingResponse{
			ResponseCode: api.RegisterWatchingResponse_FAILURE,
			Message:      "Something wrong",
			PlayerId:     "",
		}, nil
	}

	log.Println(player)

	return &api.RegisterWatchingResponse{
		ResponseCode: api.RegisterWatchingResponse_SUCCESS,
		Message:      "success",
		PlayerId:     player.ID,
	}, nil
}
