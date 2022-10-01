package dal

import (
	"context"

	"gorm.io/gorm"

	"DIdemo/entity"
)

type UserDal struct {
	DB *gorm.DB
}

func NewUserDal(db *gorm.DB) *UserDal {
	return &UserDal{
		DB: db,
	}
}

func (u *UserDal) Create(
	ctx context.Context,
	data *UserCreateParams,
) error {
	db := u.DB.Model(&entity.User{})
	// new user entity
	user := entity.User{
		Username: data.Username,
		Password: data.Password,
	}

	return db.WithContext(ctx).
		Create(&user).Error
}

// UserCreateParams for user creation
type UserCreateParams struct {
	Username string
	Password string
}
