package roman_nummerals_decoder

import (
	"strings"
)

func Decode(roman string) int {
	//Symbol    Value
	//I          1
	//V          5
	//X          10
	//L          50
	//C          100
	//D          500
	//M          1,000
	romanToArr := strings.Split(roman, "")
	var res int

	for i := 0; i < len(romanToArr); i++ {
		if romanToArr[i] == "I" {
			if romanToArr[i+1] != "I" {

			} else {
				res += 1
			}
		}
		if romanToArr[i] == "V" {
			res += 5
		}
		if romanToArr[i] == "X" {
			res += 10
		}
	}

	return res
}
