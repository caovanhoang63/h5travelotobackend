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

type checkPinPasswordBiz struct {
	cacher cacher.Cacher
}

func NewCheckPinPasswordBiz(cacher cacher.Cacher) *checkPinPasswordBiz {
	return &checkPinPasswordBiz{
		cacher: cacher,
	}
}

func (biz *checkPinPasswordBiz) CheckPinPassword(ctx context.Context, email, pin string) error {
	cachePin, err := biz.cacher.Get(ctx, common.GenResetPasswordKey(email))
	if err != nil {
		if errors.Is(err, cacher.ErrKeyNotFound) {
			return usermodel.ErrWrongPinCode
		}
		return common.ErrInternal(err)
	}
	pin = fmt.Sprintf("\"%v\"", pin)
	log.Println(pin)
	log.Println(cachePin)
	if pin != cachePin {
		err = biz.cacher.Del(ctx, common.GenResetPasswordKey(email))
		if err != nil {
			return common.ErrInternal(err)
		}
		return usermodel.ErrWrongPinCode
	}
	return nil

}
