package htsavestore

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *store) ListHotelsSaved(ctx context.Context,
	conditions map[string]interface{},
	filter *htsavemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.Hotel, error) {
	var result []htsavemodel.HotelSave
	db := s.db.Table(htsavemodel.HotelSave{}.TableName()).Where(conditions)

	//if v := filter; v != nil {
	//	if v.UserId > 0 {
	//		s.db = s.db.Where("user_id = ?", v.UserId)
	//	}
	//}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	hotels := make([]common.Hotel, len(result))
	for i, item := range result {
		result[i].Hotel.CreatedAt = item.CreatedAt
		result[i].Hotel.UpdatedAt = nil
		hotels[i] = *result[i].Hotel

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return hotels, nil
}
