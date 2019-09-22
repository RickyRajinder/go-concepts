package main

import (
	"math"
	"regexp"
	"strings"
)

var frequencyOrder= map[string]float64 {
	"A": 8.167,
	"B": 1.492,
	"C": 2.782,
	"D": 4.253,
	"E": 12.702,
	"F": 2.228,
	"G": 2.015,
	"H": 6.094,
	"I": 6.966,
	"J": 0.153,
	"K": 0.772,
	"L": 4.025,
	"M": 2.406,
	"N": 6.749,
	"O": 7.507,
	"P": 1.929,
	"Q": 0.095,
	"R": 5.987,
	"S": 6.327,
	"T": 9.056,
	"U": 2.758,
	"V": 0.978,
	"W": 2.360,
	"X": 0.15,
	"Y": 1.974,
	"Z": 0.074,
}

func getLetterCount(message string) map[rune]int {
	message = strings.ToUpper(message)
	letterCount := map[rune]int{
		'A': 0,
		'B': 0,
		'C': 0,
		'D': 0,
		'E': 0,
		'F': 0,
		'G': 0,
		'H': 0,
		'I': 0,
		'J': 0,
		'K': 0,
		'L': 0,
		'M': 0,
		'N': 0,
		'O': 0,
		'P': 0,
		'Q': 0,
		'R': 0,
		'S': 0,
		'T': 0,
		'U': 0,
		'V': 0,
		'W': 0,
		'X': 0,
		'Y': 0,
		'Z': 0,
	}
	msgArr := stringToStringArray(message)
	msgRune := []rune(message)

	for i := range msgArr {
		if strings.Contains(VALUE_SET, msgArr[i]) {
			letterCount[msgRune[i]] += 1
		}
	}
	return letterCount
}


func getFrequencyPercentage(letter string, msgArr []string) float64 {
	count := 0
	for i := range msgArr {
		if letter == strings.ToUpper(msgArr[i]) {
			count++
		}
	}
	return (float64(count) / float64(len(msgArr))) * 100
}

func getFrequencyValues(message string) map[string]float64 {
	weights := map[string]float64{}
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	message = reg.ReplaceAllString(message, "")
	msgArr := stringToStringArray(message)
	for i := range msgArr {
		weights[strings.ToUpper(msgArr[i])] = getFrequencyPercentage(strings.ToUpper(msgArr[i]), msgArr)
	}
	return weights
}

func getMinLetterWeightDifference(weight float64) string {
	minDiff := math.MaxFloat32
	letter := ""
	for k, v := range frequencyOrder {
		if minDiff > math.Abs(weight-v) {
			minDiff = math.Abs(weight - v)
			letter = k
		}
	}
	return letter
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if strings.ToUpper(v) == strings.ToUpper(str) {
			return true
		}
	}
	return false
}

func findAndReplace(message string) string {
	weights := getFrequencyValues(message)
	newString := strings.ToUpper(message)
	var seen []string
	for k, v := range weights {
		if !stringInSlice(k, seen) {
			newString = strings.ReplaceAll(newString, k, getMinLetterWeightDifference(v))
		}
		seen = append(seen, k)
	}
	return newString
}

