package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type WebErrResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
}