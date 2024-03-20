package workersqlstorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) ListHotelWorkers(ctx context.Context,
	conditions map[string]interface{},
	filter *workermodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {
	var result []workermodel.Worker
	db := s.db

	db = db.Table(workermodel.Worker{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.HotelId > 0 {
			db.Where("hotel_id = ?", v.HotelId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	fmt.Println("result", len(result))
	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
