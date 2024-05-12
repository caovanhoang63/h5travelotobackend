package rtsearchstorage

import (
	json2 "encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"strconv"
	"strings"
)

func (s *store) ListRoomTypeWithFilter(ctx context.Context,
	filter *rtsearchmodel.Filter) ([]rtsearchmodel.RoomType, error) {
	var result []rtsearchmodel.RoomType

	req, err := filter.ToSearchRequest()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	sort := types.SortOptions{
		SortOptions: map[string]types.FieldSort{
			"price": {
				Order: &sortorder.SortOrder{Name: "asc"},
			},
		},
	}

	res, err := s.es.Search().Index(rtsearchmodel.IndexName).Sort(sort).
		Request(req).Do(ctx)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	if res.Hits.Total.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var json []byte
			var roomType rtsearchmodel.RoomType
			var strFields rtsearchmodel.RoomTypeStrFields
			json, err = hit.Source_.MarshalJSON()
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &roomType)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &strFields)
			if strFields.Bed != nil {
				strings.Trim(*strFields.Bed, "\"")
				if err := json2.Unmarshal([]byte(*strFields.Bed), &roomType.Bed); err != nil {
					return nil, common.ErrInternal(err)
				}
			}

			if strFields.Images != nil {
				strings.Trim(*strFields.Images, "\"")
				if err := json2.Unmarshal([]byte(*strFields.Images), &roomType.Images); err != nil {
					return nil, common.ErrInternal(err)
				}
			}

			roomType.Id, err = strconv.Atoi(hit.Id_)

			if err != nil {
				continue
			}
			result = append(result, roomType)
		}
		return result, nil

	}

	return nil, nil
}

func (s *store) GetAvailableRoomTypeInCache(ctx context.Context,
	key string) (jsonByte []byte, err error) {

	val, err := s.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, common.ErrDb(err)
	}
	return val, nil
}
