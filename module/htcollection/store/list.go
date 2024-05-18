package htcollectionstore

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
	htcollection "h5travelotobackend/module/htcollection/model"
	"time"
)

func (s *store) ListCollectionWithCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *htcollection.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]htcollection.HotelCollection, error) {
	var result []htcollection.HotelCollection

	db := s.db.Table(htcollection.HotelCollection{}.TableName()).Where(conditions).
		Where("status = ?", common.StatusActive)

	if filter != nil {
		if filter.UserId > 0 {
			db = db.Where("user_id = ?", filter.UserId)
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

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}

func (s *store) ListHotelsInCollection(ctx context.Context,
	conditions map[string]interface{},
	filter *htsavemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.Hotel, error) {
	var result []htcollection.HotelCollectionDetail
	db := s.db.Table(htcollection.HotelCollectionDetail{}.TableName()).Where(conditions)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(common.DbTimeStampLayout, string(base58.Decode(v)))
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
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(common.DbTimeStampLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return hotels, nil
}
