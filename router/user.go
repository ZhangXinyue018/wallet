package router

import (
	"github.com/wallet/service/interf"
	"net/http"
)

type _UserRouter struct {
	UserService interf.IUserService
}

func (router *_UserRouter) GenRoutingFunc() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			router.UserService.CreateUser()
		}
	}
}
