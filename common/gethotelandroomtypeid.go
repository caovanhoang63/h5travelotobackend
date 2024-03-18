package common

import "github.com/gin-gonic/gin"

func GetHotelAndRoomTypeId(c *gin.Context) (hotelId, roomTypeId int, err error) {
	hotelUid, err := FromBase58(c.Param("hotel-id"))
	if err != nil {
		return 0, 0, err
	}
	roomTypeUid, err := FromBase58(c.Param("room-type-id"))
	if err != nil {
		return 0, 0, err
	}
	return int(hotelUid.GetLocalID()), int(roomTypeUid.GetLocalID()), nil
}
