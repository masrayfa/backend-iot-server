package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) 
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) 
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Activation(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) 
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) 
}