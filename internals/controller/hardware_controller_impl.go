package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/service"
)

type HardwareControllerImpl struct {
	service service.HardwareService	
}

func NewHardwareController(service service.HardwareService) HardwareController {
	return &HardwareControllerImpl{
		service: service,
	}
}

func (c *HardwareControllerImpl) FindHardwareTypeById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	hardwareResponse, err := c.service.FindHardwareTypeById(request.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: hardwareResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *HardwareControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hardwareResponse, err := c.service.FindAll(request.Context())
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: hardwareResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *HardwareControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	hardwareResponse, err := c.service.FindById(request.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: hardwareResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *HardwareControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hardwareRequest := web.HardwareCreateRequest{}
	helper.ReadFromRequestBody(request, &hardwareRequest)

	hardwareResponse, err := c.service.Create(request.Context(), hardwareRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: hardwareResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *HardwareControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	hardwareRequest := web.HardwareUpdateRequest{}
	helper.ReadFromRequestBody(request, &hardwareRequest)

	err = c.service.Update(request.Context(), hardwareRequest, id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *HardwareControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	err = c.service.Delete(request.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}