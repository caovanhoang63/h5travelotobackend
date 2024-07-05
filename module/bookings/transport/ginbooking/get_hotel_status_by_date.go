package ginbooking

//func GetBookingStatisticByDate(c *gin.Context) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var date *common.CivilDate
//		err := c.ShouldBindQuery(&date)
//		if err != nil {
//			panic(common.ErrInvalidRequest(err))
//		}
//		uid, err := common.FromBase58(c.Param("hotel-id"))
//		if err != nil {
//			panic(common.ErrInvalidRequest(err))
//		}
//
//		hotelId := int(uid.GetLocalID())
//
//	}
//}
