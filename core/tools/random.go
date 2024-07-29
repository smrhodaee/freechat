package tools

import (
	cRand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

const (
	digits   = "0123456789"
	specials = "~=+%^*/()[]{}/!@#$?|"
	uppers   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowers   = "abcdefghijklmnopqrstuvwxyz"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

func GenerateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := cRand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func randomString(source string, length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = source[random.Intn(len(source))]
	}
	random.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func RandomCode(length int) string {
	return randomString(digits, length)
}

func RandomChapta(length int) string {
	return randomString(uppers+digits, length)
}

func RandomToken(length int) string {
	return randomString(uppers+lowers+digits, length)
}
