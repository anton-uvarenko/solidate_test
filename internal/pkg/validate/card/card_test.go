package card

import (
	"github.com/anton-uvarenko/solidgate_test/internal/pkg/payload"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLuhnAlg(t *testing.T) {
	type testCase struct {
		CardNumber     int
		Length         int
		ExpectedOutput bool
	}

	testTable := []testCase{
		{
			CardNumber:     17893729974,
			Length:         11,
			ExpectedOutput: true,
		},
	}

	for _, v := range testTable {
		assert.Equal(t, v.ExpectedOutput, luhnAlg(v.CardNumber, v.Length))
	}
}

func TestIsCardValid(t *testing.T) {
	type testCase struct {
		Input             payload.ValidCardRequest
		ExpectedValidity  bool
		ExpectedErrorCode int
	}
	testTable := []testCase{
		{
			Input: payload.ValidCardRequest{
				CardNumber:      "371449635398431",
				ExpirationYear:  2024,
				ExpirationMonth: 12,
			},
			ExpectedValidity:  true,
			ExpectedErrorCode: 0,
		},
	}
	for _, tc := range testTable {
		validity, err := IsCardValid(tc.Input)
		assert.Equal(t, tc.ExpectedValidity, validity)

		// todo check error can be nil => panic
		if tc.ExpectedErrorCode != 0 {
			assert.Equal(t, tc.ExpectedErrorCode, err.(payload.Error).Code)
		}
	}
}
