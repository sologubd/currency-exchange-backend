package bank

type CurrencyCode string
type CountryCode string

const (
	EUR CurrencyCode = "EUR"
	PLN CurrencyCode = "PLN"
	UAH CurrencyCode = "UAH"
	USD CurrencyCode = "USD"
)

func (cc CurrencyCode) ISO_4217() int32 {
	switch cc {
	case EUR:
		return 978
	case PLN:
		return 985
	case UAH:
		return 980
	case USD:
		return 840
	default:
		return 0
	}
}

func ParseISO_4217(code int32) CurrencyCode {
	switch code {
	case 978:
		return EUR
	case 985:
		return PLN
	case 980:
		return UAH
	case 840:
		return USD
	default:
		return ""
	}
}

const (
	UA CountryCode = "UA"
	PL CountryCode = "PL"
)

type ExchangeRate struct {
	Buy  float64
	Sell float64
}

type IBank interface {
	CurrentExchangeRate() ExchangeRate
}
