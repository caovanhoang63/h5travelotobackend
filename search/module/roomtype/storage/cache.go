package rtsearchstorage

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"time"
)

func (s *store) Cache(ctx context.Context, key string, rt []rtsearchmodel.RoomType) error {
	jsonByte, err := json.Marshal(rt)
	if err != nil {
		return common.ErrInternal(err)
	}

	sttCmd := s.rdb.Set(ctx, key, jsonByte, time.Minute*5)

	if sttCmd.Err() != nil {
		return common.ErrDb(sttCmd.Err())
	}

	return nil
}
