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

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	user, err := controller.userService.FindAll(request.Context())
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: user,
	}

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("user_id")	
	userId, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	user, err := controller.userService.FindById(request.Context(), userId)
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: user,
	}

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)
	log.Println("userCreateRequest", userCreateRequest)

	userResponse, err := controller.userService.Register(request.Context(), request, userCreateRequest)
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
	token := request.URL.Query().Get("token")

	err := controller.userService.Activation(request.Context(), token)
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: nil,
	}

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	forgotPasswordRequest := web.UserForgotPasswordRequest{}
	helper.ReadFromRequestBody(request, &forgotPasswordRequest)

	err := controller.userService.ForgotPassword(request.Context(), forgotPasswordRequest)
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: nil,
	}

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateRequest := web.UserUpdatePasswordRequest{}
	helper.ReadFromRequestBody(request, &updateRequest)

	param := params.ByName("id")
	userId, err := strconv.ParseInt(param, 10, 64)
	helper.PanicIfError(err)

	err = controller.userService.MatchPassword(request.Context(), userId, updateRequest.OldPassword)
	helper.PanicIfError(err)

	err = controller.userService.UpdatePassword(request.Context(), userId, updateRequest.NewPassword)
	helper.PanicIfError(err)

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: nil,
	}

	helper.WriteToResponseBody(writer, webReponse)
}

func (controller *UserControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	err = controller.userService.Delete(request.Context(), int64(userId))

	webReponse := web.WebResponse {
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: nil,
	}

	helper.WriteToResponseBody(writer, webReponse)
}