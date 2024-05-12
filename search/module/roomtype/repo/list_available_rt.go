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
	Cache(ctx context.Context, key string, rt []rtsearchmodel.RoomType) error
}

type BookingHandler interface {
	CountBookedRoom(ctx context.Context, rtId int, startDate, endDate *common.CivilDate) (*int, error)
}

type listRoomTypeRepo struct {
	rtStore   ListRoomTypeStore
	bkHandler BookingHandler
}

func NewListRoomTypeRepo(rtStore ListRoomTypeStore, bkStore BookingHandler) *listRoomTypeRepo {
	{
		return &listRoomTypeRepo{
			rtStore:   rtStore,
			bkHandler: bkStore,
		}
	}
}

func (repo *listRoomTypeRepo) CacheRoomTypes(ctx context.Context, key string, rts []rtsearchmodel.RoomType) error {
	err := repo.rtStore.Cache(ctx, key, rts)
	if err != nil {
		return common.ErrToCacheEntity(rtsearchmodel.EntityName, err)
	}
	return nil
}

func (repo *listRoomTypeRepo) ListRoomType(ctx context.Context,
	filter *rtsearchmodel.Filter,
) ([]rtsearchmodel.RoomType, error) {
	rts, err := repo.rtStore.ListRoomTypeWithFilter(ctx, filter)
	if err != nil {
		return nil, common.ErrCannotListEntity(rtsearchmodel.EntityName, err)
	}

	if rts == nil {
		return nil, nil
	}

	var jobs []asyncjob.Job
	for i := range rts {
		if rts[i].TotalRoom == 0 {
			rts[i].AvailableRoom = 0
			continue
		}
		job := asyncjob.NewJob(func(ctx context.Context) error {

			booked, err := repo.bkHandler.CountBookedRoom(ctx, rts[i].Id, filter.StartDate, filter.EndDate)

			if err != nil {
				return common.ErrInternal(err)
			}
			if booked == nil {
				rts[i].AvailableRoom = rts[i].TotalRoom
			} else {
				rts[i].AvailableRoom = rts[i].TotalRoom - *booked
			}
			return nil

		})
		jobs = append(jobs, job)
	}

	g := asyncjob.NewGroup(true, jobs...)
	err = g.Run(ctx)
	if err != nil {
		log.Println("a: ", err)
	}

	return rts, nil
}
