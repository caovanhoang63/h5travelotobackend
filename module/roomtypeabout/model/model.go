package roomtypeaboutmodel

import (
	"errors"
	"h5travelotobackend/common"
)

const EntityName = "RoomTypeAbout"

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

type RoomTypeAboutUpdate struct {
	common.MongoModel `json:",inline" bson:",inline"`
	Convenient        map[string]string `json:"convenient,omitempty" bson:"convenient,omitempty"`
	About             string            `json:"about,omitempty" bson:"about,omitempty"`
}

func (RoomTypeAboutUpdate) CollectionName() string {
	return "room_type_about"
}

var (
	ErrInvalidRoomType = common.NewCustomError(
		errors.New("invalid room type"),
		"invalid room type",
		"INVALID_ROOM_TYPE")
	ErrRoomTypeAboutExisted = common.NewCustomError(
		errors.New("room type about existed"),
		"room type about existed",
		"ROOM_TYPE_ABOUT_EXISTED")
)
