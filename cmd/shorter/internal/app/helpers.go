package app

import (
	"math/rand"
	"strings"
	"time"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789" + "_")

func RandomShortURL() string {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rd.Intn(len(chars))])
	}
	str := b.String()
	return str
}
