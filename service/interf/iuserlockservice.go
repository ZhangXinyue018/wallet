package interf

type IUserLockService interface {
	UnlockUser()

	LockUser()

	SetUnlockTimeout()
}
