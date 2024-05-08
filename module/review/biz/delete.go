package reviewbiz

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	reviewmodel "h5travelotobackend/module/review/model"
	"log"
)

type DeleteReviewStore interface {
	Delete(ctx context.Context, Id primitive.ObjectID) error
	FindWithCondition(ctx context.Context, condition map[string]interface{}) (*reviewmodel.Review, error)
}

type deleteReviewBiz struct {
	store DeleteReviewStore
	ps    pubsub.Pubsub
}

func NewDeleteReviewBiz(store DeleteReviewStore, ps pubsub.Pubsub) *deleteReviewBiz {
	return &deleteReviewBiz{store: store}
}

func (biz *deleteReviewBiz) DeleteReview(ctx context.Context, requester common.Requester, Id primitive.ObjectID) error {
	oldData, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"_id": Id})
	if err != nil {
		return common.ErrEntityNotFound(reviewmodel.EntityName, err)
	}

	if oldData.Status == common.StatusDeleted {
		return common.ErrEntityDeleted(reviewmodel.EntityName, common.RecordNotFound)
	}

	if requester.GetRole() != common.RoleAdmin && requester.GetUserId() != oldData.UserId {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Delete(ctx, Id); err != nil {
		return common.ErrCannotDeleteEntity(reviewmodel.EntityName, err)
	}
	mess := pubsub.NewMessage(common.DTOReview{
		HotelId: oldData.HotelId,
		UserId:  oldData.UserId,
		Rating:  oldData.Rating,
	})

	if err := biz.ps.Publish(ctx, common.TopicUserDeleteReviewHotel, mess); err != nil {
		log.Println(common.ErrCannotPublishMessage(common.TopicUserDeleteReviewHotel, err))
	}

	return nil
}
