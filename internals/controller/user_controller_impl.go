package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/service"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)
	log.Println("userCreateRequest", userCreateRequest)

	userResponse, err := controller.userService.Register(request.Context(), userCreateRequest)
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: userResponse,
	}	

	log.Println("webResponse", userResponse)

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	userResponse, err := controller.userService.Login(request.Context(), loginRequest)
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) Activation(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *UserControllerImpl) ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *UserControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}