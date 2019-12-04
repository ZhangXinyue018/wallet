package router

import (
	"github.com/wallet/service/interf"
	"net/http"
)

type _SignatureRouter struct {
	SignatureService interf.ISignatureService
}

func (router *_SignatureRouter) GenRoutingFunc() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			router.SignatureService.Sign()
		case http.MethodGet:
			router.SignatureService.CheckSig()
		}
	}
}
