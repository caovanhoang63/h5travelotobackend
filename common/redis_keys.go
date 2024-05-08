package common

import "fmt"

func GetRateLimitKey(clientIp string, api string) string {
	return fmt.Sprintf("RATE_LIMIT_COUNT_%s_%s", clientIp, api)
}
