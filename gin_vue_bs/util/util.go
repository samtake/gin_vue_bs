package util

import (
	"math/rand"
	"time"
)

//返回随机字符串
func RandomString(n int) string {
	var letters = []byte("iasdhjklfhascvxnjklasdfhjkasdfklasdfhnjklasdfjklasdfjklfasdsdfjk")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
