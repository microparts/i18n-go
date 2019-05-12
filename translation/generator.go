package translation

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func RandomStr() string {
	return hex.EncodeToString(
		sha256.New().
			Sum([]byte(
				strconv.Itoa(10000000 + rand.Intn(100000000)),
			)))
}

func GenerateString(languages []string) String {
	str := (&String{}).Init()

	for _, lang := range languages {
		str.Translate[lang] = RandomStr()
	}

	return *str
}
