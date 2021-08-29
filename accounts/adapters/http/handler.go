package http

type HttpResponse struct {
	Body  interface{}
	Code  uint16
	Error error
}

type HttpHandler interface {
	Handle(request interface{}) (response HttpResponse)
}
