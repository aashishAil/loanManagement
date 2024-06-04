package constant

type Currency string

const (
	CurrencyINR                 Currency = "INR"
	MinCurrencyConversionFactor          = 100
)

var (
	CurrencyMap = map[Currency]struct{}{
		CurrencyINR: {},
	}
)
