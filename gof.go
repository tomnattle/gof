package gof

type Gof struct {
}

type IRequest interface {
}
type IResponse interface {
}

type RequestHandler func(IRequest, IResponse)

func (g *Gof) Router(path string, handler RequestHandler) {

}

func (g *Gof) Start() {

}
