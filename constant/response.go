package constant

type Response struct {
	Error string `json:"error"`
}

var (
	DefaultErrorResponse = Response{
		Error: "something went wrong",
	}
	EmptyEmailResponse = Response{
		Error: "email cannot be empty",
	}
	EmptyPasswordResponse = Response{
		Error: "password cannot be empty",
	}
	InvalidInputResponse = Response{
		Error: "invalid input",
	}
	MissingAuthTokenResponse = Response{
		Error: "missing auth token",
	}
	UnableToAuthenticateResponse = Response{
		Error: "unable to authenticate",
	}
)
