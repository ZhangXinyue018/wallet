package router

import (
	"net/http"
)

var (
	UserRouter      IRouter
	UserLockRouter  IRouter
	UserKeyRouter   IRouter
	SignatureRouter IRouter
)

func init() {
	UserRouter = &_UserRouter{}
	UserLockRouter = &_UserLockRouter{}
	UserKeyRouter = &_UserKeyRouter{}
	SignatureRouter = &_SignatureRouter{}
}

type IRouter interface {
	GenRoutingFunc() func(w http.ResponseWriter, req *http.Request)
}
