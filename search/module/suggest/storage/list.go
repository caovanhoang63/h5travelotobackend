package suggeststorage

import (
	json2 "encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	suggestmodel "h5travelotobackend/search/module/suggest/model"
	"strings"
)

func (s *esStore) ListSuggestions(ctx context.Context,
	input *suggestmodel.SuggestRequest,
) (*suggestmodel.SuggestResponse, error) {
	var response suggestmodel.SuggestResponse
	var result suggestmodel.SuggestionHits

	boost1 := map[string]types.Float64{
		"hotels_enriched": 28.0,
	}
	boost2 := map[string]types.Float64{
		"provinces": 30.0,
	}
	boost3 := map[string]types.Float64{
		"districts": 10.0,
	}
	boot4 := map[string]types.Float64{
		"wards": 10.0,
	}
	boost5 := map[string]types.Float64{
		"landmarks_enriched": 90.0,
	}

	str := strings.Split(input.SearchText, " ")
	wildCard := "*" + str[len(str)-1] + "*"
	caseInsensitive := true

	req := &search.Request{
		IndicesBoost: []map[string]types.Float64{boost1, boost2, boost3, boot4, boost5},
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Should: []types.Query{
					{
						Match: map[string]types.MatchQuery{
							"name": {Query: input.SearchText},
						},
					},
					{
						Wildcard: map[string]types.WildcardQuery{
							"name": {
								Value:           &wildCard,
								CaseInsensitive: &caseInsensitive,
							},
						},
					},
				},
			},
		},
	}

	do := s.es.Search().
		Index("hotels_enriched,provinces,districts,wards,landmarks_enriched").
		Request(req).
		Size(input.Limit)

	res, err := do.Do(ctx)

	if err != nil {
		return nil, common.ErrDb(err)
	}

	if res.Hits.Total.Value > 0 {
		response.Total = res.Hits.Total.Value
		var json []byte
		for _, hit := range res.Hits.Hits {
			var suggest suggestmodel.SuggestionHit
			json, err = hit.Source_.MarshalJSON()
			if err != nil {
				return nil, err
			}
			err = json2.Unmarshal(json, &suggest)
			if err != nil {
				return nil, err
			}
			suggest.Score = hit.Score_
			suggest.Id = hit.Id_
			suggest.Index = hit.Index_
			result.Append(suggest)
		}

	}

	response.SuggestionHits = &result

	return &response, nil
}
