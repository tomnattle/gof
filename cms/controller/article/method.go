package article

import (
	"context"

	"github.com/tomnattle/gof"
)

type Article struct {
}

func (p *Article) Create(*context.Context, *gof.Request, *gof.Response) (err error) { return }
func (p *Article) Delete(*context.Context, *gof.Request, *gof.Response) (err error) { return }
func (p *Article) Update(*context.Context, *gof.Request, *gof.Response) (err error) { return }
func (p *Article) Get(*context.Context, *gof.Request, *gof.Response) (err error)    { return }
func (p *Article) List(*context.Context, *gof.Request, *gof.Response) (err error)   { return }
