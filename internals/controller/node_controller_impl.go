package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

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
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	log.Println("limit: ", limit)

	node, err := controller.nodeService.FindAll(request.Context(), limit)
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
		Data:   node,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	limitStr := request.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	startDateStr := request.URL.Query().Get("start")
	var startDate *time.Time 
	if startDateStr == "" {
		start, err := time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			webErrResponse := web.WebErrResponse{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Mesage: err.Error(),
			}

			helper.WriteToResponseBody(writer, webErrResponse)
			return
		}
		startDate = &start
	}

	endDateStr := request.URL.Query().Get("end")
	var endDate *time.Time
	
	if endDateStr == "" {
		end, err := time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			webErrResponse := web.WebErrResponse{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Mesage: err.Error(),
			}

			helper.WriteToResponseBody(writer, webErrResponse)
			return
		}

		endDate = &end
	}

	// startDateStr := request.URL.Query().Get("start")
	// var startDate int64

	// if startDateStr == "" {
	// 	startDateStr = "0"
	// } else {
	// 	startDate, err = strconv.ParseInt(startDateStr, 10, 64)
	// 	if err != nil {
	// 		webErrResponse := web.WebErrResponse{
	// 			Code: http.StatusBadRequest,
	// 			Status: http.StatusText(http.StatusBadRequest),
	// 			Mesage: err.Error(),
	// 		}

	// 		helper.WriteToResponseBody(writer, webErrResponse)
	// 		return
	// 	}
	// }

	// endDateStr := request.URL.Query().Get("end")
	// var endDate int64
	
	// if endDateStr == "" {
	// 	endDateStr = "0"

	// } else {
	// 	endDate, err = strconv.ParseInt(endDateStr, 10, 64)
	// 	if err != nil {
	// 		webErrResponse := web.WebErrResponse{
	// 			Code: http.StatusBadRequest,
	// 			Status: http.StatusText(http.StatusBadRequest),
	// 			Mesage: err.Error(),
	// 		}

	// 		helper.WriteToResponseBody(writer, webErrResponse)
	// 		return
	// 	}
	// }

	node, err := controller.nodeService.FindById(request.Context(), id, limit, startDate, endDate)
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
		Data:   node,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	nodeCreateRequest := web.NodeCreateRequest{}
	helper.ReadFromRequestBody(request, &nodeCreateRequest)

	_, err := controller.nodeService.Create(request.Context(), nodeCreateRequest)
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

func (controller *NodeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	nodeUpdateRequest := web.NodeUpdateRequest{}
	helper.ReadFromRequestBody(request, &nodeUpdateRequest)

	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	err = controller.nodeService.Update(request.Context(), nodeUpdateRequest, id)
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

	log.Println("Node with id: ", id, " has been updated")

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	err = controller.nodeService.Delete(request.Context(), id)
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
		Data:   nil,
	}

	log.Println("Node with id: ", id, " has been deleted")

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NodeControllerImpl) FindHardwareNode(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	hardware, err := controller.nodeService.FindHardwareNode(request.Context(), id)
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
		Data:   hardware,
	}

	helper.WriteToResponseBody(writer, webResponse)
}