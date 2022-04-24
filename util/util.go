package util

import (
	"math/rand"
	"time"
)

func RandomString(i int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, i)

	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
