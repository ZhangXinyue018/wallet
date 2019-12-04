package interf

type IUserKeyService interface {
	ImportKeys()

	CreateKey()

	ListKeys()
}
