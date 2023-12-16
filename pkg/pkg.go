package pkg

type App struct {
	Port string
}

type HttpError struct {
	Message string
	Code    int
}

func (e *HttpError) Error() string {
	return e.Message
}
