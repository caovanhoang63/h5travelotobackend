package ginhotelsearch

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelsearchbiz "h5travelotobackend/search/module/hotel/biz"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	hotelstorage "h5travelotobackend/search/module/hotel/storage/esstore"
	"log"
	"net/http"
)

// url: GET v1/hotels/search
// params:
//	type Filter struct {
//	SearchText   string     `json:"search_text"`
//	Adults       int        `json:"adults"`
//	Children     int        `json:"children"`
//	Star         int        `json:"star"`
//	ByLocation   bool       `json:"by_location"`
//	ByProvince   bool       `json:"by_province"`
//	ByDistrict   bool       `json:"by_district"`
//	ByWard       bool       `json:"by_ward"`
//	ByHotelName  bool       `json:"by_hotel_name"`
//	ProvinceCode string     `json:"province_code"`
//	DistrictCode string     `json:"district_code"`
//	WardCode     string     `json:"ward_code"`
//	Name         string     `json:"name"`
//	Lat          float64    `json:"lat"`
//	Lng          float64    `json:"lng"`
//	StartDate    *time.Time `json:"start_date"`
//	EndDate      *time.Time `json:"date_date"`
//}
//

func ListHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter hotelmodel.Filter
		err := c.ShouldBindQuery(&filter)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		log.Println("filter: ", filter)

		var paging common.Paging
		err = c.ShouldBindQuery(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelstorage.NewESStore(appCtx.GetElasticSearchClient())
		biz := hotelsearchbiz.ListHotelStore(store)
		result, err := biz.ListHotel(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, filter, paging))
	}
}
