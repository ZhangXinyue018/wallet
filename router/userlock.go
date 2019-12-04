package router

import (
	"github.com/wallet/service/interf"
	"net/http"
)

type _UserLockRouter struct {
	UserLockService interf.IUserLockService
}

func (router *_UserLockRouter) GenRoutingFunc() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			router.UserLockService.LockUser()
		case http.MethodDelete:
			router.UserLockService.UnlockUser()
		case http.MethodPut:
			router.UserLockService.SetUnlockTimeout()
		}
	}
}
