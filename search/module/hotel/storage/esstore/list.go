package hotelstorage

//func (s *esStore) ListHotel(ctx context.Context, filter *hotelmodel.Filter, paging *common.Paging) ([]hotelmodel.Hotel, error) {
//
//	queryJson, err := json.Marshal(filter)
//
//	if err != nil {
//		return nil, err
//	}
//
//	from := paging.GetOffSet()
//	size := paging.Limit
//
//	s.es.Search().Index(hotelmodel.IndexName).Request(&search.Request{Query: })
//
//}
