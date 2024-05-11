package rtsearchstorage

import (
	json2 "encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"strconv"
)

func (s *store) ListRoomTypeWithFilter(ctx context.Context,
	filter *rtsearchmodel.Filter) ([]rtsearchmodel.RoomType, error) {
	var result []rtsearchmodel.RoomType

	req, err := filter.ToSearchRequest()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	res, err := s.es.Search().Index(rtsearchmodel.IndexName).
		Request(req).Do(ctx)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	if res.Hits.Total.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var json []byte
			var roomType rtsearchmodel.RoomType
			err = hit.Source_.UnmarshalJSON(json)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &roomType)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			roomType.Id, err = strconv.Atoi(hit.Id_)
			if err != nil {
				continue
			}
			result = append(result, roomType)
		}
	}

	return result, nil
}
