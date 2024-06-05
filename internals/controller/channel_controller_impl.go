package controller

import (
	"log"
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

	log.Println("channelRequest: ", channelRequest)

	_, err := controller.channelService.Create(request.Context(), channelRequest)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, webResponse)
}