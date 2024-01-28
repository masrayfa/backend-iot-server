package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/service"
)

type ChannelControllerImpl struct {
	channelService service.ChannelService
}

func NewChannelController(channelService service.ChannelService) ChannelController {
	return &ChannelControllerImpl{
		channelService: channelService,
	}
}

func (controller *ChannelControllerImpl) GetNodeChannel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (controller *ChannelControllerImpl) GetNodeChannelMultiple(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (controller *ChannelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}