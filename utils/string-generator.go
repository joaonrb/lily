//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package utils

import (
	"math/rand"
	"time"
)

const symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits # This case 63 = 0b111111
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var random = rand.NewSource(time.Now().UnixNano())

// Solution based in http://stackoverflow.com/users/1705598/icza solution at
// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func GenerateBase64Bytes(n int) []byte {
	result := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, random.Int63(), letterIdxMax; i >= 0; i-- {
		if remain == 0 {
			cache, remain = random.Int63(), letterIdxMax
		}
		result[i] = symbols[int(cache & letterIdxMask)]  // Cache & letter = the first 6 bits of cache (0-63 rand)

		cache >>= letterIdxBits  // Shift 6 bits
		remain--
	}
	return result
}

func GenerateBase64String(n int) string {
	return string(GenerateBase64Bytes(n))
}
