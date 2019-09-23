package main

import (
	"math/rand"
	"strings"
	"text/scanner"
	"time"
	"unicode"
)

const VALUE_SET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encrypt(key, message string) string {
	return translateMessage(key, message, "E")
}

func decrypt(key, message string) string {
	return translateMessage(key, message, "D")
}

func stringToStringArray(str string) []string {
	vals := []rune(str)
	var res []string

	for _, char := range vals {
		res = append(res, strings.Replace(scanner.TokenString(char), "\"", "", -1))
	}
	return res
}

//Generates a (pseudo) random permutation of the alphabet that serves as our key
func generateKey() string {
	res := stringToStringArray(VALUE_SET)
	final := make([]string, len(res))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(res))

	for i, j := range perm {
		final[j] = res[i]
	}

	strKey := strings.Join(final, "")
	return strKey
}

//Depending on `mode`, will decrypt or encrypt the given message
//by resolving it against the key
func translateMessage(key, message, mode string) string {
	translatedMsg := ""
	charsA := VALUE_SET
	charsB := key
	if mode == "D" {
		charsA, charsB = charsB, charsA
	}
	msgArr := stringToStringArray(message)
	charsBArr := stringToStringArray(charsB)
	var index int
	for i := range msgArr {
		currentSymbol := msgArr[i]
		if strings.Contains(charsA, strings.ToUpper(currentSymbol)) {
			index = strings.Index(charsA, strings.ToUpper(currentSymbol))
		}
		msgRune := []rune(message)
		if unicode.IsUpper(msgRune[i]) {
			translatedMsg += strings.ToUpper(charsBArr[index])
		} else if unicode.IsLower(msgRune[i]) {
			translatedMsg += strings.ToLower(charsBArr[index])
		} else {
			translatedMsg += currentSymbol
		}
	}
	return translatedMsg
}

