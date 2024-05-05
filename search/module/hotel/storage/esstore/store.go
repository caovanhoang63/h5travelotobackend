package hotelstorage

import "github.com/elastic/go-elasticsearch/v8"

type esStore struct {
	es *elasticsearch.TypedClient
}

func NewESStore(es *elasticsearch.TypedClient) *esStore {
	return &esStore{es}
}
