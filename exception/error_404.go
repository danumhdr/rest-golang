package exception

type Error404 struct {
	Error string
}

func NewError404(err string) Error404 {
	return Error404{Error: err}
}
