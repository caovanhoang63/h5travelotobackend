package common

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type Location struct {
	Lat *types.Float64 `json:"lat"`
	Lon *types.Float64 `json:"lon"`
}
