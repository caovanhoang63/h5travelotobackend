package reviewrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

type ListReviewsStore interface {
	ListReviewWithCondition(
		ctx context.Context,
		filter *reviewmodel.Filter,
		paging *common.Paging,
		cond map[string]interface{},
		moreKeys ...string,
	) ([]reviewmodel.Review, error)
}

type ListUserStore interface {
	ListUsersByIds(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

type listReviewsRepo struct {
	store     ListReviewsStore
	userStore ListUserStore
}

func NewListReviewsRepo(store ListReviewsStore, userStore ListUserStore) *listReviewsRepo {
	return &listReviewsRepo{store: store, userStore: userStore}
}

func (repo *listReviewsRepo) ListReviewsWithCondition(
	ctx context.Context,
	filter *reviewmodel.Filter,
	paging *common.Paging,
) ([]reviewmodel.Review, error) {
	data, err := repo.store.ListReviewWithCondition(ctx, filter, paging, nil)
	if err != nil {
		return nil, common.ErrCannotListEntity(reviewmodel.EntityName, err)
	}

	ids := make([]int, len(data))
	for i := range data {
		ids[i] = data[i].UserId
	}

	users, err := repo.userStore.ListUsersByIds(ctx, ids)

	if err != nil {
		return nil, common.ErrCannotListEntity(common.SimpleUser{}.TableName(), err)
	}

	userMap := make(map[int]*common.SimpleUser)
	for i := range users {
		user := users[i]
		userMap[users[i].Id] = &user
	}

	for i := range data {
		data[i].User = userMap[data[i].UserId]
	}

	return data, nil
}
