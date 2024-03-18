package roomtypeaboutmodel

import (
	"errors"
	"h5travelotobackend/common"
)

const EntityName = "RoomType"

type RoomTypeAbout struct {
	common.MongoModel `json:",inline" bson:",inline"`
	RoomTypeId        int               `json:"-" bson:"room_type_id"`
	FakeRoomTypeId    string            `json:"room_type_id" bson:"fake_room_type_id"`
	Convenient        map[string]string `json:"convenient" bson:"convenient"`
	About             string            `json:"about" bson:"about"`
}

func (RoomTypeAbout) CollectionName() string {
	return "room_type_about"
}

func (r *RoomTypeAbout) Validate() error {
	return nil
}

var (
	ErrInvalidRoomType = common.NewCustomError(
		errors.New("invalid room type"),
		"invalid room type",
		"INVALID_ROOM_TYPE")
)
