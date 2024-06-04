package constant

type Response struct {
	Error string `json:"error"`
}

var (
	AdminOnlyRouteResponse = Response{
		Error: "user should be an admin",
	}
	CustomerOnlyRouteResponse = Response{
		Error: "user should be a customer",
	}
	DefaultErrorResponse = Response{
		Error: "something went wrong",
	}
	EmptyEmailResponse = Response{
		Error: "email cannot be empty",
	}
	EmptyPasswordResponse = Response{
		Error: "password cannot be empty",
	}
	InvalidAmountResponse = Response{
		Error: "invalid amount",
	}
	InvalidCurrencyResponse = Response{
		Error: "currency cannot be empty",
	}
	InvalidFetchScheduledRepaymentsResponse = Response{
		Error: "invalid flag value for fetchScheduledRepayments",
	}
	InvalidInputResponse = Response{
		Error: "invalid input",
	}
	InvalidStatusResponse = Response{
		Error: "invalid status",
	}
	InvalidTermResponse = Response{
		Error: "invalid term",
	}
	InvalidDisbursalDateResponse = Response{
		Error: "invalid disbursal date",
	}
	MissingAuthTokenResponse = Response{
		Error: "missing auth token",
	}
	UnableToAuthenticateResponse = Response{
		Error: "unable to authenticate",
	}
)
