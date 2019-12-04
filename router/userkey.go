package router

import (
	"github.com/wallet/service/interf"
	"net/http"
)

type _UserKeyRouter struct {
	UserKeyService interf.IUserKeyService
}

func (router *_UserKeyRouter) GenRoutingFunc() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			router.UserKeyService.CreateKey()
			//router.UserKeyService.ImportKeys()
		case http.MethodGet:
			router.UserKeyService.ListKeys()

		}
	}
}
