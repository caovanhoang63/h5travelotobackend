package hotelstorage

import (
	json2 "encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

func (s *esStore) ListHotel(ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging,
) ([]hotelmodel.Hotel, error) {
	var result []hotelmodel.Hotel

	req, err := filter.ToSearchRequest()
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	res, err := s.es.Search().Index(hotelmodel.IndexName).
		Request(req).Do(ctx)
	if err != nil {
		return nil, err
	}

	if res.Hits.Total.Value > 0 {
		var json []byte
		for _, hit := range res.Hits.Hits {
			var hotel hotelmodel.Hotel
			json, err = hit.Source_.MarshalJSON()
			if err != nil {
				return nil, err
			}
			err = json2.Unmarshal(json, &hotel)
			if err != nil {
				return nil, err
			}
			hotel.Id = hit.Id_
			result = append(result, hotel)
		}

	} else {
		fmt.Print("Found no documents\n")
	}

	return result, err
}
