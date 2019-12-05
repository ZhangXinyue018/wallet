package router

import (
	"net/http"
)

var (
	UserRouter      IRouter
	UserLoginRouter IRouter
	UserKeyRouter   IRouter
	SignatureRouter IRouter
)

func init() {
	UserRouter = &_UserRouter{}
	UserLoginRouter = &_UserLoginRouter{}
	UserKeyRouter = &_UserKeyRouter{}
	SignatureRouter = &_SignatureRouter{}
}

type IRouter interface {
	GenRoutingFunc() func(w http.ResponseWriter, req *http.Request)
}
