package gof

import "context"

type Gof struct {
}

func NewGof() *Gof {
	return new(Gof)
}

type Request interface {
}
type Response interface {
}

type RequestHandler func(context.Context, Request, Response) (err error)

func (g *Gof) Router(path string, handler RequestHandler) {

}

func (g *Gof) Start() {

}
