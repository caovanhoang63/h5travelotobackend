package reviewbiz

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

type listReviewsBiz struct {
	store ListReviewsStore
}

func NewListReviewsBiz(store ListReviewsStore) *listReviewsBiz {
	return &listReviewsBiz{store: store}
}

func (biz *listReviewsBiz) ListReviews(
	ctx context.Context,
	filter *reviewmodel.Filter,
	paging *common.Paging,
) ([]reviewmodel.Review, error) {
	data, err := biz.store.ListReviewWithCondition(ctx, filter, paging, nil)
	if err != nil {
		return nil, common.ErrCannotListEntity(reviewmodel.EntityName, err)
	}
	return data, nil
}
