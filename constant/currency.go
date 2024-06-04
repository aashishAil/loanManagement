package constant

type Currency string

const (
	CurrencyINR Currency = "INR"
)

var (
	CurrencyMap = map[Currency]struct{}{
		CurrencyINR: {},
	}
)
