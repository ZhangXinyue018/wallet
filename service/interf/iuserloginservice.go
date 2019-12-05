package interf

type IUserLoginService interface {
	Login(user, password string) string

	Logout(user, password string)
}
