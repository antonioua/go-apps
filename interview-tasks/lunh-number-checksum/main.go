package main

import (
	"strconv"
	"strings"
)

func luhn(cardNum string) bool {

	charList := strings.Split(cardNum, "")

	sumDouble := 0
	sumOther := 0

	for i := len(charList) - 1; i >= 0; i-- {
		parsedDigit, err := strconv.Atoi(charList[i])
		if err != nil {
			panic("Error parsing digit")
		}

		sumOther = sumOther + parsedDigit

		if i == 0 {
			break
		}

		i--
		parsedDigit, err = strconv.Atoi(charList[i])
		if err != nil {
			panic("Error parsing digit")
		}

		parsedDigit = parsedDigit * 2

		if parsedDigit >= 10 {
			sumDouble += parsedDigit - 9
		} else {
			sumDouble += parsedDigit
		}
	}

	reSum := sumOther + sumDouble

	if (reSum)%10 != 0 {
		return false
	}

	return true
}
