package rtsearchbiz

import (
	"golang.org/x/net/context"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
)

type ListAvailableRtRepo interface {
	ListRoomType(ctx context.Context,
		filter *rtsearchmodel.Filter,
	) ([]rtsearchmodel.RoomType, error)
}

type listAvailableRtBiz struct {
	rtRepo ListAvailableRtRepo
}

func NewListAvailableRtBiz(rtRepo ListAvailableRtRepo) *listAvailableRtBiz {
	return &listAvailableRtBiz{rtRepo: rtRepo}
}

func (biz *listAvailableRtBiz) ListAvailableRt(ctx context.Context,
	filter *rtsearchmodel.Filter,
) ([]rtsearchmodel.RoomType, error) {
	rts, err := biz.rtRepo.ListRoomType(ctx, filter)
	if err != nil {
		return nil, err
	}
	for i := len(rts) - 1; i >= 0; i-- {
		if rts[i].AvailableRoom == 0 || rts[i].AvailableRoom < filter.RoomQuantity ||
			float32(rts[i].MaxCustomer*filter.RoomQuantity) < filter.Customer {
			rts = append(rts[:i], rts[i+1:]...)
		}
	}

	if len(rts) == 0 {
		return nil, nil
	}

	return rts, nil
}
