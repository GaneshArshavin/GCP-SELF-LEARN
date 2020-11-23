package redis

import "fmt"

func GetloginAttemptsKey(username string) string {
	return fmt.Sprintf("user_%s", username)
}

func GetInHouseTokenKey(token string) string {
	return fmt.Sprintf("inhouse_%s", token)
}

func GetThirdPartyTokenKey(token string, companyName string) string {
	return fmt.Sprintf("inhouse_%s_%s", token)
}
