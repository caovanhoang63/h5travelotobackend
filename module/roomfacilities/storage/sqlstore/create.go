package roomfacilitysqlstore

import (
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roomfacilitymodel "h5travelotobackend/module/roomfacilities/model"
)

func (s *sqlStore) CreateHotelFacilityDetails(
	ctx context.Context,
	facilities []roomfacilitymodel.RoomFacilityDetail) error {
	fmt.Println("IdIDIDIDID", facilities[1].FacilityId)
	if err := s.db.Create(facilities).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
