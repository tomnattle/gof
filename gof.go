package gof

import "context"

type Gof struct {
}

func New() *Gof {
	return new(Gof)
}

type Request interface {
}

type Response interface {
}

type RequestHandler func(*context.Context, *Request, *Response) (err error)

type RestRequestHandler interface {
	List(*context.Context, *Request, *Response) (err error)
	Get(*context.Context, *Request, *Response) (err error)
	Create(*context.Context, *Request, *Response) (err error)
	Delete(*context.Context, *Request, *Response) (err error)
	Update(*context.Context, *Request, *Response) (err error)
}

func (g *Gof) Router(path string, handler RequestHandler) {

}

func (g *Gof) Get(path string, handler RequestHandler) {

}

func (g *Gof) Post(path string, handler RequestHandler) {

}

func (g *Gof) Rest(path string, handler RestRequestHandler) {

}

func (g *Gof) Start() {

}
