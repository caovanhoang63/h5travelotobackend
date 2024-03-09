package userbiz

import (
	"context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

type UserUpdateStore interface {
	Update(ctx context.Context, id int, update *usermodel.UserUpdate) error
}

type userUpdateBiz struct {
	store UserUpdateStore
}

func NewUserUpdateBiz(store UserUpdateStore) *userUpdateBiz {
	return &userUpdateBiz{store: store}
}

func (biz *userUpdateBiz) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
