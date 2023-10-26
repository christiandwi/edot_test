package guest

import (
	"encoding/base64"
	"log"

	"github.com/christiandwi/edot/user-service/domain/users"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo users.UsersRepository
}

func NewGuestService(userRepo users.UsersRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (serv *service) LoginService(userIdentifier string, password string) (token string, err error) {
	userData, err := serv.userRepo.GetUser(userIdentifier)
	if err != nil {
		log.Print("error at getting user")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error at hashing password, err : %s", err.Error())
		return
	}

	if bcrypt.CompareHashAndPassword(hashed, []byte(userData.Password)) != nil {
		log.Print("invalid password")
		return
	}

	token = base64.StdEncoding.EncodeToString([]byte(userData.SecureId))
	return
}
