package ginpayin

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/payment/vnpay"
	"log"
)

func VnpIPN(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var ipn vnpay.IPNResponse

		err := c.ShouldBind(&ipn)
		if err != nil {
			log.Println("Error binding IPN response: ", err)
		}
		log.Println("IPN response: ", ipn)

	}
}
