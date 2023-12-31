package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorPassword       = 10007

	ErrorDatabase = 40001
	ErrorCreate   = 40002
	ErrorCheck    = 40003
	ErrorList     = 40003
	ErrorSearch   = 40005
	ErrorUpdate   = 40006
	ErrorDelete   = 40007

	ErrorAuthCheckTokenFail    = 30001
	ErrorAuthCheckTokenTimeout = 30002
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorId                    = 30005
)
