package suggestbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	suggestmodel "h5travelotobackend/search/module/suggest/model"
)

type ListSuggestStore interface {
	ListSuggestions(ctx context.Context,
		input *suggestmodel.SuggestRequest,
	) (*suggestmodel.SuggestResponse, error)
}

type listSuggestBiz struct {
	store ListSuggestStore
}

func NewListSuggestBiz(store ListSuggestStore) *listSuggestBiz {
	return &listSuggestBiz{store: store}
}

func (biz *listSuggestBiz) ListSuggestions(ctx context.Context,
	input *suggestmodel.SuggestRequest,
) (*suggestmodel.SuggestResponse, error) {

	response, err := biz.store.ListSuggestions(ctx, input)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return response, nil
}
