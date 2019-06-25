package gof

import "context"

type Gof struct {
}

type Request interface {
}
type Response interface {
}

type RequestHandler func(context.Context, Request, Response)

func (g *Gof) Router(path string, handler RequestHandler) {

}

func (g *Gof) Start() {

}
