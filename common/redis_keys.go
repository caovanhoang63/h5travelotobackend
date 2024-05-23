package common

import "fmt"

// =================== Cache ===========================

const CacheKey = "cache_key"

func GenApiCacheKey(url string) string {
	return fmt.Sprintf("cache:%s", url)
}

func GenKeyForDelApiCache(entity string, id string) string {
	return fmt.Sprintf("cache:v/%s/%s", entity, id)
}

// ==================== Cache ===========================

// =================== Rate limiting ===========================

func GenRateLimitKeyById(id int, api string) string {
	return fmt.Sprintf("rate_limit_count:%v:%s", id, api)
}

func GenRateLimitKeyByIp(ip, api string) string {
	return fmt.Sprintf("rate_limit_count:%v:%s", ip, api)
}

func GenBanUserIdKey(id int) string {
	return fmt.Sprintf("ban_user:%v", id)
}

func GenBanUserIpKey(ip string) string {
	return fmt.Sprintf("ban_user:%v", ip)
}

// ==================== Rate limiting ===========================

func GenUserRecentViewedKey(id int) string {
	return fmt.Sprintf("recent_viewed:%v", id)
}

func GenResetPasswordKey(email string) string {
	return fmt.Sprintf("reset:%s", email)
}
