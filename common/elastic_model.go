package common

type ElasticModel struct {
	Id    string  `json:"_id"`
	Score float64 `json:"_score"`
}
