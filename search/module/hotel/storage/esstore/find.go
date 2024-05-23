package hotelstorage

import (
	json2 "encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	"strconv"
)

func (s *esStore) FindHotelById(ctx context.Context, id int) (*hotelmodel.Hotel, error) {
	hit, err := s.es.Get(hotelmodel.IndexName, strconv.Itoa(id)).Do(ctx)
	if err != nil {
		return nil, common.ErrDb(err)
	}
	var hotel hotelmodel.Hotel
	jsonData, err := hit.Source_.MarshalJSON()
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	err = json2.Unmarshal(jsonData, &hotel)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	hotel.Id, err = strconv.Atoi(hit.Id_)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	hotel.Mask(false)

	return &hotel, nil
}
