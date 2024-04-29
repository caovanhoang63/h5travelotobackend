package userbiz

import (
	"context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

type UserChangePasswordStore interface {
	ChangePassword(ctx context.Context, id int, data *usermodel.UserChangePassword) error
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type userChangePasswordBiz struct {
	store  UserChangePasswordStore
	hasher Hasher
}

func NewUserChangePasswordBiz(store UserChangePasswordStore, hasher Hasher) *userChangePasswordBiz {
	return &userChangePasswordBiz{store: store, hasher: hasher}
}

func (biz *userChangePasswordBiz) ChangePassword(ctx context.Context, id int, data *usermodel.UserChangePassword) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if biz.hasher.Hash(data.OldPassword+user.Salt) != user.Password {
		return usermodel.ErrOldPasswordNotMatch
	}

	if biz.hasher.Hash(data.Password+user.Salt) == user.Password {
		return usermodel.ErrPasswordMatchedWithPast
	}

	if err := biz.store.ChangePassword(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
