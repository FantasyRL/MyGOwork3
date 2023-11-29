package model

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type TokenResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Token  string      `json:"token"`
	Expire string      `json:"expire"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
