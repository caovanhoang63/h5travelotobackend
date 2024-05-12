package utils

import "fmt"

func GenCacheKeyForQuery(queryTime int64, hotelId int) string {
	return fmt.Sprintf("%v:hotel:%v", queryTime, hotelId)
}
