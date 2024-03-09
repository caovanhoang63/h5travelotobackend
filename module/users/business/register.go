package userbiz

import (
	"context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(registerStorage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		if user.Status == 0 {
			return usermodel.ErrUserHasDeletedOrDisabled
		}
		return usermodel.ErrEmailExisted
	}

	data.Salt = common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + data.Salt)
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
