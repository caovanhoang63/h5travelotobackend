package reviewstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *mongoStore) FindWithCondition(ctx context.Context, condition map[string]interface{}) (*reviewmodel.Review, error) {
	coll := s.db.Collection(reviewmodel.Review{}.CollectionName())

	result := coll.FindOne(ctx, condition)
	if err := result.Err(); err != nil {
		return nil, common.ErrDb(err)
	}

	var data reviewmodel.Review
	if err := result.Decode(&data); err != nil {
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
