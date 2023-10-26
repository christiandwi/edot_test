package users

import "github.com/christiandwi/edot/user-service/entity"

type UsersRepository interface {
	GetUser(userIdentifier string) (user entity.Users, err error)
}
