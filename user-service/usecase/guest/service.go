package guest

type Service interface {
	LoginService(username string, password string) (token string, err error)
}
