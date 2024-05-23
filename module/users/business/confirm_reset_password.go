package userbiz

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/cacher"
	"h5travelotobackend/email"
	usermodel "h5travelotobackend/module/users/model"
	"math/rand"
	"time"
)

type ConfirmResetPasswordBiz struct {
	store      CheckExistedEmailStorage
	cacher     cacher.Cacher
	mailSender email.Engine
}

func NewConfirmResetPassword(store CheckExistedEmailStorage, cacher cacher.Cacher, mailSender email.Engine) *ConfirmResetPasswordBiz {
	return &ConfirmResetPasswordBiz{
		store:      store,
		cacher:     cacher,
		mailSender: mailSender,
	}
}

func (biz *ConfirmResetPasswordBiz) ConfirmResetPassword(ctx context.Context, userEmail string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": userEmail})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return usermodel.ErrInvalidEmail
		}
		return common.ErrInternal(err)
	}

	if user.Status != 1 {
		return usermodel.ErrUserBanned
	}

	pin := NewPinCode()
	err = biz.cacher.Set(ctx, common.GenResetPasswordKey(userEmail), pin, time.Minute*5)
	if err != nil {
		return common.ErrInternal(err)
	}
	mail := email.NewRecoverPasswordMail(userEmail, pin)

	biz.mailSender.Send(*mail)

	return nil
}

func NewPinCode() string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	one := random.Intn(10)
	two := random.Intn(10)
	three := random.Intn(10)
	four := random.Intn(10)

	pin := fmt.Sprintf("%d%d%d%d", one, two, three, four)

	return pin
}
