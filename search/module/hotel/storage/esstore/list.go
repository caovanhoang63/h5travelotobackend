package hotelstorage

import (
	json2 "encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	"strconv"
	"strings"
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
		Request(req).From(paging.GetOffSet()).Size(paging.Limit).Do(ctx)
	if err != nil {
		return nil, err
	}

	if res.Hits.Total.Value > 0 {
		paging.Total = res.Hits.Total.Value
		var json []byte
		for _, hit := range res.Hits.Hits {
			var hotel hotelmodel.Hotel
			var hotelImage hotelmodel.HotelImage
			json, err = hit.Source_.MarshalJSON()
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &hotel)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &hotelImage)
			if hotelImage.LogoStr != nil {
				strings.Trim(*hotelImage.LogoStr, "\"")
				if err = json2.Unmarshal([]byte(*hotelImage.LogoStr), &hotel.Logo); err != nil {
					return nil, common.ErrInternal(err)
				}
			}
			if hotelImage.ImagesStr != nil {
				strings.Trim(*hotelImage.ImagesStr, "\"")
				if err = json2.Unmarshal([]byte(*hotelImage.ImagesStr), &hotel.Images); err != nil {
					return nil, common.ErrInternal(err)
				}
			}

			hotel.Id, err = strconv.Atoi(hit.Id_)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			hotel.Mask(false)
			result = append(result, hotel)
		}

	} else {
		fmt.Print("Found no documents\n")
	}

	return result, err
}

func (s *esStore) ListRandomHotels(ctx context.Context, limit int) ([]hotelmodel.Hotel, error) {
	var result []hotelmodel.Hotel
	res, err := s.es.Search().Index(hotelmodel.IndexName).
		Request(&search.Request{
			Query: &types.Query{
				FunctionScore: &types.FunctionScoreQuery{
					Functions: []types.FunctionScore{
						{
							RandomScore: &types.RandomScoreFunction{},
						},
					},
					Query: &types.Query{
						MatchAll: &types.MatchAllQuery{},
					},
				},
			}}).
		Size(limit).Do(ctx)

	if err != nil {
		return nil, err
	}

	if res.Hits.Total.Value > 0 {
		var json []byte
		for _, hit := range res.Hits.Hits {
			var hotel hotelmodel.Hotel
			var hotelImage hotelmodel.HotelImage
			json, err = hit.Source_.MarshalJSON()
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &hotel)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			err = json2.Unmarshal(json, &hotelImage)
			if hotelImage.LogoStr != nil {
				strings.Trim(*hotelImage.LogoStr, "\"")
				if err = json2.Unmarshal([]byte(*hotelImage.LogoStr), &hotel.Logo); err != nil {
					return nil, common.ErrInternal(err)
				}
			}
			if hotelImage.ImagesStr != nil {
				strings.Trim(*hotelImage.ImagesStr, "\"")
				if err = json2.Unmarshal([]byte(*hotelImage.ImagesStr), &hotel.Images); err != nil {
					return nil, common.ErrInternal(err)
				}
			}
			hotel.Id, err = strconv.Atoi(hit.Id_)
			if err != nil {
				return nil, common.ErrInternal(err)
			}
			hotel.Mask(false)
			result = append(result, hotel)
		}
	} else {
		fmt.Print("Found no documents\n")
	}

	return result, err
}

func (s *esStore) ListByIds(ctx context.Context, ids []int) ([]hotelmodel.Hotel, error) {
	var result []hotelmodel.Hotel
	var idsStr []string
	for _, i := range ids {

		idsStr = append(idsStr, strconv.Itoa(i))
	}
	res, err := s.es.Mget().Index(hotelmodel.IndexName).Ids(idsStr...).Do(ctx)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	for _, doc := range res.Docs {
		hit := doc.(*types.GetResult)
		var json []byte
		var hotel hotelmodel.Hotel
		var hotelImage hotelmodel.HotelImage
		json, err = hit.Source_.MarshalJSON()
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		err = json2.Unmarshal(json, &hotel)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		if hotel.Status == 0 {
			continue
		}
		err = json2.Unmarshal(json, &hotelImage)
		if hotelImage.LogoStr != nil {
			strings.Trim(*hotelImage.LogoStr, "\"")
			if err = json2.Unmarshal([]byte(*hotelImage.LogoStr), &hotel.Logo); err != nil {
				return nil, common.ErrInternal(err)
			}
		}
		if hotelImage.ImagesStr != nil {
			strings.Trim(*hotelImage.ImagesStr, "\"")
			if err = json2.Unmarshal([]byte(*hotelImage.ImagesStr), &hotel.Images); err != nil {
				return nil, common.ErrInternal(err)
			}
		}
		hotel.Id, err = strconv.Atoi(hit.Id_)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		hotel.Mask(false)
		result = append(result, hotel)
	}
	return result, nil
}
