package rtsearchrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/asyncjob"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"log"
)

type ListRoomTypeStore interface {
	ListRoomTypeWithFilter(ctx context.Context,
		filter *rtsearchmodel.Filter) ([]rtsearchmodel.RoomType, error)
}

type BookingStore interface {
	CountBookedRoom(ctx context.Context, rtId int, startDate, endDate *common.CivilDate) (*int, error)
}

type listRoomTypeRepo struct {
	rtStore ListRoomTypeStore
	bkStore BookingStore
}

func NewListRoomTypeRepo(rtStore ListRoomTypeStore, bkStore BookingStore) *listRoomTypeRepo {
	{
		return &listRoomTypeRepo{
			rtStore: rtStore,
			bkStore: bkStore,
		}
	}
}

func (repo *listRoomTypeRepo) ListRoomType(ctx context.Context,
	filter *rtsearchmodel.Filter,
) ([]rtsearchmodel.RoomType, error) {
	rts, err := repo.rtStore.ListRoomTypeWithFilter(ctx, filter)
	if err != nil {
		return nil, common.ErrCannotListEntity(rtsearchmodel.EntityName, err)
	}

	jobs := make([]asyncjob.Job, len(rts))
	for i := range rts {
		asyncjob.NewJob(func(ctx context.Context) error {
			var booked *int
			booked, err = repo.bkStore.CountBookedRoom(ctx, rts[i].Id, filter.StartDate, filter.EndDate)
			if err != nil || booked == nil {
				return common.ErrInternal(err)
			}
			rts[i].AvailableRoom = rts[i].TotalRoom - *booked
			return nil
		})
		jobs = append(jobs, jobs[i])
	}

	g := asyncjob.NewGroup(true, jobs...)
	err = g.Run(ctx)
	if err != nil {
		log.Println("a: ", err)
	}

	return rts, nil
}
