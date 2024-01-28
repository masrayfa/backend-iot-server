package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ChannelController interface {
	GetNodeChannel(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetNodeChannelMultiple(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}