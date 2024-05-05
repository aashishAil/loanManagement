package constant

type Response struct {
	Error string `json:"error"`
}

var (
	InvalidInputResponse = Response{
		Error: "invalid input",
	}
	DefaultErrorResponse = Response{
		Error: "something went wrong",
	}
)
