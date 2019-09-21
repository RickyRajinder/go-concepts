package main

import (
	"math/rand"
	"strings"
	"text/scanner"
	"time"
)

const VALUE_SET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateKey() string {
	vals := []rune(VALUE_SET)
	var res []string

	for _, char := range vals {
		res = append(res, scanner.TokenString(char))
	}

	final := make([]string, len(res))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(res))

	for i, j := range perm {
		final[j] = res[i]
	}

	strKey := strings.Join(final, "")
	strKey = strings.Replace(strKey, "\"", "", -1)
	return strKey
}

func main() {
	print(generateKey())
}

