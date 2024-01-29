package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/web"
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

func (controller *ChannelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	channelRequest := web.ChannelCreateRequest{}
	helper.ReadFromRequestBody(request, &channelRequest)

	channelResponse, err := controller.channelService.Create(request.Context(), channelRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   channelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}