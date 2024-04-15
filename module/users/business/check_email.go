package userbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

type CheckExistedEmailStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type CheckExistedEmailBiz struct {
	store CheckExistedEmailStorage
}

func NewCheckExistedEmailBiz(store CheckExistedEmailStorage) *CheckExistedEmailBiz {
	return &CheckExistedEmailBiz{store: store}
}

func (biz *CheckExistedEmailBiz) CheckExistedEmail(ctx context.Context, email string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})

	if err != nil {
		if err == common.RecordNotFound {
			return nil
		}
		return common.ErrInternal(err)
	}

	if user != nil {
		if user.Status == 0 {
			return usermodel.ErrUserHasDeletedOrDisabled
		}
		return usermodel.ErrEmailExisted
	}

	return nil
}
