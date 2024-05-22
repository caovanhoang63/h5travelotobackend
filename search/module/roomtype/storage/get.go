package rtsearchstorage

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
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
