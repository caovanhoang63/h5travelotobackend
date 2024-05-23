package userbiz

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/cacher"
	usermodel "h5travelotobackend/module/users/model"
	"log"
)

type changePasswordForgotBiz struct {
	store  UserChangePasswordStore
	cacher cacher.Cacher
	hasher Hasher
}

func NewChangePasswordForgot(store UserChangePasswordStore, cacher cacher.Cacher, hasher Hasher) *changePasswordForgotBiz {
	return &changePasswordForgotBiz{
		store:  store,
		cacher: cacher,
		hasher: hasher,
	}
}

func (biz *changePasswordForgotBiz) ChangePassword(ctx context.Context,
	userMail,
	pin string,
	data *usermodel.UserChangePassword) error {

	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": userMail})
	if err != nil {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	log.Println(common.GenResetPasswordKey(userMail))
	cachePin, err := biz.cacher.Get(ctx, common.GenResetPasswordKey(userMail))
	if err != nil {
		if errors.Is(err, cacher.ErrKeyNotFound) {
			return usermodel.ErrWrongPinCode
		}
		return common.ErrInternal(err)
	}

	log.Println("pin:", pin)
	log.Println("cacher pin:", cachePin)
	pin = fmt.Sprintf("\"%v\"", pin)

	if pin != cachePin {
		return usermodel.ErrWrongPinCode
	}

	hashed := biz.hasher.Hash(data.Password + user.Salt)
	if hashed == user.Password {
		return usermodel.ErrPasswordMatchedWithPast
	}

	if err = data.Validate(); err != nil {
		return err
	}

	data.Password = hashed
	if err = biz.store.ChangePassword(ctx, user.Id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
