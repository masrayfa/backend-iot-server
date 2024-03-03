package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/service"
)

type NodeControllerImpl struct {
	nodeService service.NodeService
}

func NewNodeController(nodeService service.NodeService) NodeController {
	return &NodeControllerImpl{nodeService: nodeService}
}

func (controller *NodeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	limitStr := request.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	helper.PanicIfError(err)

	log.Println("limit: ", limit)

	node, err := controller.nodeService.FindAll(request.Context(), limit, 1)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   node,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	limitStr := request.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	helper.PanicIfError(err)

	node, err := controller.nodeService.FindById(request.Context(), id, limit)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   node,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	nodeCreateRequest := web.NodeCreateRequest{}
	helper.ReadFromRequestBody(request, &nodeCreateRequest)

	nodeCreateResponse, err := controller.nodeService.Create(request.Context(), nodeCreateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nodeCreateResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	nodeUpdateRequest := web.NodeUpdateRequest{}
	helper.ReadFromRequestBody(request, &nodeUpdateRequest)

	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	err = controller.nodeService.Update(request.Context(), nodeUpdateRequest, id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	log.Println("Node with id: ", id, " has been updated")

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	err = controller.nodeService.Delete(request.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	log.Println("Node with id: ", id, " has been deleted")

	helper.WriteToResponseBody(writer, webResponse)
}