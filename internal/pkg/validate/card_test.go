package validate

import (
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

	assert.Equal(t, testTable[0].ExpectedOutput, luhnAlg(testTable[0].CardNumber, testTable[0].Length))
}
