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

//type TokenData struct {
//	User  User   `json:"user"`
//	Token string `json:"token"`
//}
