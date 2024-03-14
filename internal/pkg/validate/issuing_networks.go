package validate

type IssueNetwork struct {
	NetworkName string
	Length      []Range
	IINRanges   []Range
	Validation  ValidationName
}

type Range struct {
	From int
	To   int
}

type ValidationName string

const (
	LuhnAlgorithmValidation ValidationName = "LuhnAlgorithm"
	NoValidation            ValidationName = "NoValidation"
)

var IssuedNetworks = []IssueNetwork{
	{
		NetworkName: "American Express",
		Length: []Range{
			{From: 15, To: 15},
		},
		IINRanges: []Range{
			{From: 34, To: 34},
			{From: 37, To: 37},
		},
		Validation: LuhnAlgorithmValidation,
	},
	{
		NetworkName: "China T-Union",
		Length: []Range{
			{From: 19, To: 19},
		},
		IINRanges: []Range{
			{From: 31, To: 31},
		},
	},
	{
		NetworkName: "Diners Club International",
		Length: []Range{
			{From: 14, To: 19},
		},
		IINRanges: []Range{
			{From: 36, To: 36},
		},
		Validation: LuhnAlgorithmValidation,
	},
	{
		NetworkName: "UkrCard",
		Length: []Range{
			{From: 16, To: 19},
		},
		IINRanges: []Range{
			{From: 60400100, To: 60420099},
		},
		Validation: LuhnAlgorithmValidation,
	},
	{
		NetworkName: "Visa",
		Length: []Range{
			{From: 13, To: 13},
			{From: 16, To: 16},
			{From: 19, To: 19},
		},
		IINRanges: []Range{
			{From: 4, To: 4},
		},
		Validation: LuhnAlgorithmValidation,
	},
}
