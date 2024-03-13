package hotelmysqlstorage

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) ListHotelWithCondition(
	ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]hotelmodel.Hotel, error) {
	var result []hotelmodel.Hotel

	db := s.db.Table(hotelmodel.Hotel{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}

		//TODO: Tinh khoang cach dua vao kinh do va vi do
		if f.Distance > 0 {

			db = db.Where("ACOS( SIN(lat)*SIN(?) + COS(lat)*COS(?)*COS(?-lng) ) * 6371000 < ?",
				f.Lat, f.Lat, f.Lng, f.Distance)
		}

		if f.DistrictID > 0 {
			db = db.Where("district_id = ?", f.DistrictID)
		}
		if f.WardID > 0 {
			db = db.Where("ward_id = ?", f.WardID)
		}
		if f.ProvinceID > 0 {
			db = db.Where("province_id = ?", f.ProvinceID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	// paging
	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("id  < ? ", uid.GetLocalID())
	} else {
		db = db.Offset(paging.GetOffSet())
	}

	if err := db.Limit(paging.Limit).Find(&result).Order("id desc").Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
