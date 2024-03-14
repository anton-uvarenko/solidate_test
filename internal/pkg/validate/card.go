package validate

import (
	"fmt"
	"github.com/anton-uvarenko/solidgate_test/internal/pkg/payload"
	"github.com/anton-uvarenko/solidgate_test/pkg/math"
	"strconv"
	"time"
)

func IsCardValid(request payload.ValidCardRequest) (bool, error) {
	expirationDate, err := time.Parse(
		time.DateOnly,
		fmt.Sprintf("%d-%d-01", request.ExpirationYear, request.ExpirationMonth),
	)
	if err != nil {
		return false, ErrInvalidDate
	}

	if expirationDate.Before(time.Now()) {
		return false, ErrInvalidDate
	}

	cardNumber, err := strconv.Atoi(request.CardNumber)
	if err != nil {
		return false, err
	}

	if !luhnAlg(cardNumber, len(request.CardNumber)) {
		return false, ErrCheckSum
	}

	if !isIssuingNetworkDataValid(request.CardNumber) {
		return false, ErrIssuingNetworkData
	}

	return true, nil
}

func isIssuingNetworkDataValid(cardNumber string) bool {
	for _, data := range IssuedNetworks {
		if isFirstDigitsValid(cardNumber, data.IINRanges) && isLengthValid(cardNumber, data.Length) {
			return true
		}
	}

	return false
}

func isFirstDigitsValid(cardNumber string, IInRanges []Range) bool {
	for _, v := range IInRanges {
		n := len(strconv.Itoa(v.From))
		firstNDigits, _ := strconv.Atoi(cardNumber[:n])
		if firstNDigits >= v.From && firstNDigits <= v.To {
			return true
		}
	}

	return false
}

func isLengthValid(cardNumber string, lengthRanges []Range) bool {
	for _, ln := range lengthRanges {
		if len(cardNumber) >= ln.From && len(cardNumber) <= ln.To {
			return true
		}
	}

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
