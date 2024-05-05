package appError

type Custom struct {
	Err  error
	Code int
}

func (c Custom) Error() string {
	return c.Err.Error()
}
