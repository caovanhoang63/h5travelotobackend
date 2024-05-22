package suggestmodel

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"h5travelotobackend/common"
)

const IndexName = "hotels_enriched,provinces,districts,wards,landmarks_enriched"

type SuggestRequest struct {
	SearchText string `json:"search_text" form:"search_text" binding:"required"`
	Limit      int    `json:"limit" form:"limit"`
}

type SuggestResponse struct {
	SuggestionHits *SuggestionHits `json:"hits"`
	Total          int64           `json:"total"`
}

type SuggestionHits []SuggestionHit

func (hits *SuggestionHits) Append(hit SuggestionHit) {
	*hits = append(*hits, hit)
}

type SuggestionHit struct {
	Index    string                     `json:"index"`
	Id       string                     `json:"id"`
	Name     string                     `json:"name"`
	FullName *string                    `json:"full_name"`
	Score    types.Float64              `json:"score"`
	Location *common.Location           `json:"location"`
	Province *common.AdministrativeUnit `json:"province"`
	District *common.AdministrativeUnit `json:"district"`
}
