package hotelstorage

import (
	json2 "encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	"log"
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
		paging.Total = res.Hits.Total.Value
		var json []byte
		for _, hit := range res.Hits.Hits {
			var hotel hotelmodel.Hotel
			//log.Println("hit: ", hit.Source_)
			json, err = hit.Source_.MarshalJSON()
			if err != nil {
				return nil, err
			}
			log.Println("json: ", string(json))
			err = json2.Unmarshal(json, &hotel)
			if err != nil {
				return nil, err
			}

			result = append(result, hotel)
		}

	} else {
		fmt.Print("Found no documents\n")
	}

	return result, err
}
