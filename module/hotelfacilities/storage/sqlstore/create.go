package hotelfacilitysqlstore

import (
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

func (s *sqlStore) CreateHotelFacilityDetails(
	ctx context.Context,
	facilities []hotelfacilitymodel.HotelFacilityDetail) error {
	fmt.Println("IdIDIDIDID", facilities[1].FacilityId)
	if err := s.db.Create(facilities).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
