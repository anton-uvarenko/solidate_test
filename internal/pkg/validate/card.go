package validate

import (
	"fmt"
	"solidgate_test/internal/pkg/payload"
	"solidgate_test/pkg/math"
	"strconv"
)

func IsCardValid(request payload.ValidCardRequest) {

}
func isFirstDigitValid(c rune) bool {
	firstDigit, err := strconv.Atoi(string(c))
	if err != nil {
		return false
	}

	fmt.Println(firstDigit)
	return false
}

func luhnAlg(number int, length int) bool {
	// remove the most right digit
	checkDigit := number % 10
	number = (number - checkDigit) / 10

	sum := 0
	for i := range length - 1 {
		if i%2 == 0 {
			doubled := (number % 10) * 2
			sum += math.SumOf(doubled)
			number = (number - number%10) / 10
			continue
		}

		sum += number % 10
		number = (number - number%10) / 10
	}

	checkDigitOverAlg := 10 - sum%10
	if checkDigitOverAlg == checkDigit {
		return true
	}

	return false
}
