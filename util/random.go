package util

import (
	"math/rand"
	"time"
)

var randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int) int {
	return min + randSource.Intn(max-min)
}

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[randSource.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[randSource.Intn(n)]
}

