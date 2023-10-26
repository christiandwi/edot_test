package repo

import (
	"github.com/christiandwi/edot/user-service/database"
	"github.com/christiandwi/edot/user-service/domain/users"
	"github.com/christiandwi/edot/user-service/entity"
)

type userRepo struct {
	db *database.Database
}

func SetupUserRepo(db *database.Database) users.UsersRepository {
	return &userRepo{db: db}
}

func (u userRepo) GetUser(userIdentifier string) (user entity.Users, err error) {
	err = u.db.
		Where("phone_number = ?", userIdentifier).
		Or("email = ?", userIdentifier).
		First(&user).Error

	return
}
