package common

import "fmt"

func GetRateLimitKey(clientIp string, api string) string {
	return fmt.Sprintf("rate_limit_count:%s:%s", clientIp, api)
}

func GetApiCacheKey(url string) string {
	return fmt.Sprintf("cache:%s", url)
}
