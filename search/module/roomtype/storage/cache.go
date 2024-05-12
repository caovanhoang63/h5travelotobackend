package rtsearchstorage

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"time"
)

func (s *store) Cache(ctx context.Context, key string, rts []rtsearchmodel.RoomType, filter *rtsearchmodel.Filter) error {
	data := common.NewSuccessResponse(rts, nil, filter)
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return common.ErrInternal(err)
	}

	if err := s.rdb.Set(ctx, key, jsonByte, time.Minute*5).Err(); err != nil {
		return common.ErrDb(err)
	}

	return nil
}
