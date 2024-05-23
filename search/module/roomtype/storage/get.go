package rtsearchstorage

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"strconv"
)

func (s *store) GetMinPriceByHotelId(ctx context.Context, hotelId int) (float64, error) {
	price := "price"
	aggName := "min_price"

	res, err := s.es.Search().Index(rtsearchmodel.IndexName).Size(0).
		Request(&search.Request{
			Aggregations: map[string]types.Aggregations{
				aggName: {
					Min: &types.MinAggregation{
						Field: &price,
					},
				},
			},
			Query: &types.Query{
				Term: map[string]types.TermQuery{
					"hotel_id": {
						Value: hotelId,
					},
				},
			},
		}).Do(ctx)
	if err != nil {
		return 0, err
	}

	agg := res.Aggregations[aggName]
	minAgg := agg.(*types.MinAggregate)
	val := minAgg.Value
	float, err := common.EsFloatToGoFloat(val)
	if err != nil {
		return 0, err
	}

	return float, nil
}

func (s *store) GetRoomTypeById(ctx context.Context, id int) (*rtsearchmodel.RoomType, error) {
	hit, err := s.es.Get(rtsearchmodel.IndexName, strconv.Itoa(id)).Do(ctx)
	if err != nil {
		return nil, common.ErrDb(err)
	}
	var roomType rtsearchmodel.RoomType
	jsonData, err := hit.Source_.MarshalJSON()
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	err = json.Unmarshal(jsonData, &roomType)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	roomType.Id, err = strconv.Atoi(hit.Id_)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	roomType.Mask(false)
	return &roomType, nil
}
