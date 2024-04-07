package reviewbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	reviewmodel "h5travelotobackend/module/review/model"
)

type CreateReviewStore interface {
	Create(ctx context.Context, review *reviewmodel.Review) error
}

type createReviewBiz struct {
	store CreateReviewStore
	ps    pubsub.Pubsub
}

func NewCreateReviewBiz(store CreateReviewStore, ps pubsub.Pubsub) *createReviewBiz {
	return &createReviewBiz{store: store, ps: ps}
}

func (biz *createReviewBiz) CreateReview(ctx context.Context, review *reviewmodel.Review) error {
	review.OnCreate()
	if err := biz.store.Create(ctx, review); err != nil {
		return common.ErrCannotCreateEntity(reviewmodel.EntityName, err)
	}
	return nil
}
