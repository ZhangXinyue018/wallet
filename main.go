package wallet

import (
	"github.com/wallet/router"
	"net/http"
)

func main() {

}

func routing() {
	http.HandleFunc("/user/", router.UserRouter.GenRoutingFunc())

	http.HandleFunc("/user-login/", router.UserLoginRouter.GenRoutingFunc())

	http.HandleFunc("/user-key", router.UserKeyRouter.GenRoutingFunc())

	http.HandleFunc("/sign", router.SignatureRouter.GenRoutingFunc())
}
