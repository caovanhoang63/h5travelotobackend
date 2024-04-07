package hotelstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *hotelmodel.HotelUpdate) error {
	if err := s.db.Table(update.TableName()).Where("id = ?", id).Updates(update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) UpdateReviewWhenUserReview(ctx context.Context, review *common.DTOReview) error {
	db := s.db.Table(hotelmodel.Hotel{}.TableName()).Where("id = ?", review.HotelId)

	db = db.Update("rating", gorm.Expr("(rating * total_rating + ?)/(total_rating + ?)", review.Rating, 1))
	db = db.Update("total_rating", gorm.Expr("total_rating + ?", 1))

	if err := db.Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

func (s *sqlStore) UpdateTotalRoomTypeAndAvgPrice(ctx context.Context, room *common.DTORoomType) error {
	db := s.db.Table(hotelmodel.Hotel{}.TableName()).Where("id = ?", room.HotelId)

	db = db.Update("avg_price", gorm.Expr("(avg_price * total_room_type + ?)/(total_room_type + ?)", room.Price, 1))
	db = db.Update("total_room_type", gorm.Expr("total_room_type + ?", 1))

	if err := db.Error; err != nil {
		return common.ErrDb(err)
	}

	return nil

}
