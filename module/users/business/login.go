package userbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/tokenprovider"
	usermodel "h5travelotobackend/module/users/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	appCtx        appContext.AppContext
	loginStorage  LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	accessExpiry  int
	refreshExpiry int
}

func NewLoginBiz(appCtx appContext.AppContext, loginStorage LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, accessExpiry, refreshExpiry int) *loginBiz {
	return &loginBiz{
		appCtx:        appCtx,
		loginStorage:  loginStorage,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.loginStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	//biz.tokenConfig.GetAtExp() ===> biz.accessExpiry

	accessToken, err := biz.tokenProvider.Generate(payload, biz.accessExpiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.refreshExpiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
