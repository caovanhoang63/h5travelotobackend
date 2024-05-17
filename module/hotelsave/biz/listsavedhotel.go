package htsavebiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

type ListHotelSavedStore interface {
	ListHotelsSaved(ctx context.Context,
		conditions map[string]interface{},
		filter *htsavemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.Hotel, error)
}

type listHotelSavedBiz struct {
	store ListHotelSavedStore
}

func NewListHotelSavedBiz(store ListHotelSavedStore) *listHotelSavedBiz {
	return &listHotelSavedBiz{
		store: store,
	}
}

func (biz *listHotelSavedBiz) ListHotelsSaveByUser(ctx context.Context,
	filter *htsavemodel.Filter,
	paging *common.Paging,
	requester common.Requester) ([]common.Hotel, error) {

	cond := map[string]interface{}{
		"user_id": requester.GetUserId(),
	}

	result, err := biz.store.ListHotelsSaved(ctx, cond,
		nil, paging,
		"Hotel", "Hotel.Province",
		"Hotel.District", "Hotel.Ward")

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return result, nil
}
