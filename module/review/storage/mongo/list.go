package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *store) ListReviewWithCondition(
	ctx context.Context,
	filter *reviewmodel.Filter,
	paging *common.Paging,
	cond map[string]interface{},
	moreKeys ...string,
) ([]reviewmodel.Review, error) {
	var result []reviewmodel.Review
	filterB, _ := filter.ToBsonD()
	filterB = append(filterB, bson.E{Key: "status", Value: 1})

	option := &options.FindOptions{}

	if v := paging.FakeCursor; v == "" {
		option.SetSkip(int64(paging.GetOffSet()))
	} else {
		cursor := primitive.ObjectID{}
		err := cursor.UnmarshalText([]byte(v))
		if err != nil {
			filterB = append(filterB, bson.E{Key: "_id", Value: bson.M{"$lt": cursor}})
		} else {
			option.SetSkip(int64(paging.GetOffSet()))
		}
	}

	option.SetLimit(int64(paging.Limit)).SetSort(bson.D{{Key: "_id", Value: -1}})

	coll := s.db.Collection(reviewmodel.Review{}.CollectionName())
	if count, err := coll.CountDocuments(ctx, filterB); err != nil {
		return nil, common.ErrDb(err)
	} else {
		paging.Total = count
	}

	cur, err := coll.Find(ctx, filterB, option)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	for cur.Next(ctx) {
		var review reviewmodel.Review
		if err := cur.Decode(&review); err != nil {
			return nil, common.ErrDb(err)
		}
		result = append(result, review)
	}

	if err := cur.Err(); err != nil {
		return nil, common.ErrDb(err)
	}

	if err := cur.Close(ctx); err != nil {
		return nil, common.ErrDb(err)
	}

	if len(result) > 0 {
		paging.NextCursor = result[len(result)-1].ID.Hex()
	}

	return result, nil
}
