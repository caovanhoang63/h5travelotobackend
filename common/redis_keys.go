package common

import "fmt"

const CacheKey = "cache_key"

func GenRateLimitKey(clientIp string, api string) string {
	return fmt.Sprintf("rate_limit_count:%s:%s", clientIp, api)
}

func GenApiCacheKey(url string) string {
	return fmt.Sprintf("cache:%s", url)
}

func GenKeyForDelApiCache(entity string, id string) string {
	return fmt.Sprintf("cache:v/%s/%s", entity, id)
}
