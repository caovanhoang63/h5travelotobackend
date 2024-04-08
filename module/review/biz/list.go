package reviewbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

type ListReviewsRepo interface {
	ListReviewsWithCondition(
		ctx context.Context,
		filter *reviewmodel.Filter,
		paging *common.Paging,
	) ([]reviewmodel.Review, error)
}

type listReviewsBiz struct {
	repo ListReviewsRepo
}

func NewListReviewsBiz(repo ListReviewsRepo) *listReviewsBiz {
	return &listReviewsBiz{repo: repo}
}

func (biz *listReviewsBiz) ListReviews(
	ctx context.Context,
	filter *reviewmodel.Filter,
	paging *common.Paging,
) ([]reviewmodel.Review, error) {
	var data []reviewmodel.Review

	data, err := biz.repo.ListReviewsWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(reviewmodel.EntityName, err)
	}

	for i := range data {
		if data[i].User.Status == common.StatusDeleted && data[i].Status == common.StatusDeleted {
			data = append(data[:i], data[i+1:]...)
			i--
		}
	}

	return data, nil
}
