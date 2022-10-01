package service

import (
	"context"

	"DIdemo/dal"
	"DIdemo/schema"
)

type UserService struct {
	UserDal *dal.UserDal
}

func NewUserService(
	userDal dal.UserDal,
) *UserService {
	return &UserService{
		UserDal: &userDal, // changed
	}
}

func (u *UserService) Register(
	ctx context.Context,
	data *schema.RegisterReq,
) (*schema.RegisterRes, error) {

	params := dal.UserCreateParams{
		Username: data.Username,
		Password: data.Password,
	}

	err := u.UserDal.Create(ctx, &params)
	if err != nil {
		return nil, err
	}

	registerRes := schema.RegisterRes{
		Msg: "register success",
	}

	return &registerRes, nil
}
