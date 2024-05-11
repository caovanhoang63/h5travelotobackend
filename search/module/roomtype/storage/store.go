package rtsearchstorage

import "github.com/elastic/go-elasticsearch/v8"

type store struct {
	es *elasticsearch.TypedClient
}

func NewStore(es *elasticsearch.TypedClient) *store {
	return &store{es: es}
}
