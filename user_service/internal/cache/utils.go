package cache

import (
	"fmt"
	"time"
)

const (
	ExpiredTime        = time.Second * 3
	SmsString   string = "123456"
)

// phoneKey && smsKey
func (u *userCache) redisForVerificationKey(verificationKey string) (string, string) {
	phoneKey := fmt.Sprintf("phone_%s", verificationKey)
	smsKey := fmt.Sprintf("sms_%s", verificationKey)
	return phoneKey, smsKey
}
