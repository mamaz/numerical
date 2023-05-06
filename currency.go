package numerical

import (
	"strconv"
	"strings"
)

// ConvertToCurrency convert a number to currency format
func ConvertToCurrency(number uint, separator string) (string, error) {
	strNumber := strconv.Itoa(int(number))
	strNumLength := len(strNumber)

	if strNumLength < 4 {
		return strNumber, nil
	}

	remaining := strNumLength % 3
	var sb strings.Builder

	if remaining == 0 {
		for index, num := range strNumber {
			if remaining == 0 && index < (strNumLength-1) && (index+1)%3 == 0 {
				sb.WriteString(string(num))
				sb.WriteString(separator)
				continue
			}

			sb.WriteString(string(num))
		}
	} else {
		front := strNumber[0:remaining]

		for index, num := range front {
			if index == remaining-1 {
				sb.WriteString(string(num))
				sb.WriteString(separator)
				continue
			}
			sb.WriteString(string(num))
		}

		back := strNumber[remaining:]

		for index, num := range back {
			if index < (len(back)-1) && (index+1)%3 == 0 {
				sb.WriteString(string(num))
				sb.WriteString(separator)
				continue
			}

			sb.WriteString(string(num))
		}
	}

	return sb.String(), nil
}
