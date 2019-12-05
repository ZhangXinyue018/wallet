package router

import (
	"github.com/wallet/service/interf"
	"net/http"
)

type _UserLoginRouter struct {
	UserLoginService interf.IUserLoginService
}

func (router *_UserLoginRouter) GenRoutingFunc() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			router.UserLoginService.Login("", "")
		case http.MethodDelete:
			router.UserLoginService.Logout("", "")
		}
	}
}
