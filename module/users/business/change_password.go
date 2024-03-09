package userbiz

import (
	"context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

type UserChangePasswordStore interface {
	ChangePassword(ctx context.Context, id int, data *usermodel.UserChangePassword) error
}

type userChangePasswordBiz struct {
	store UserChangePasswordStore
}

func NewUserChangePasswordBiz(store UserChangePasswordStore) *userChangePasswordBiz {
	return &userChangePasswordBiz{store: store}
}

func (biz *userChangePasswordBiz) ChangePassword(ctx context.Context, id int, data *usermodel.UserChangePassword) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.ChangePassword(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
