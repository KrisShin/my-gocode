package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26)
		cha := r.Intn(2)
		if cha == 1 {
			b += 65 // 大写字母
		} else {
			b += 97 // 小写字母
		}
		bytes[i] = byte(b)
	}
	return string(bytes)
}
