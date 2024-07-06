package userbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/tokenprovider"
	usermodel "h5travelotobackend/module/users/model"
)

func (biz *loginBiz) Renew(ctx context.Context, token *tokenprovider.Token) (*tokenprovider.Token, error) {
	payload, err := biz.tokenProvider.Validate(token.Token)

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	user, err := biz.loginStorage.FindUser(ctx, map[string]interface{}{"id": payload.UserId})
	if err != nil {
		return nil, common.ErrInternal(err)
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}
	}

	if user.Status == common.StatusDeleted {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if user.Status != common.StatusActive {
		return nil, usermodel.ErrUserBanned
	}

	newToken, err := biz.tokenProvider.Generate(*payload, biz.accessExpiry)

	if user.Role == common.RoleOwner ||
		user.Role == common.RoleManager ||
		user.Role == common.RoleStaff {
		worker, err := biz.workerStorage.FindWithCondition(ctx, map[string]interface{}{
			"user_id": user.Id,
		})
		if err == nil && worker != nil {

		}
	}

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return newToken, nil

}
