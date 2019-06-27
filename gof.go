package gof

import (
	"context"
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"

	"github.com/tomnattle/gof/module/log"
)

type Gof struct {
	Env    string
	Logger log.Logger
	Config Conf
}

func New() *Gof {
	return new(Gof)
}
func (p *Gof) SetEnv(env string) {
	p.Env = env
}

type Request interface {
}

type Response interface {
}
type ControllerConfig struct {
	Denies []reflect.Type
}

type RequestHandler func(*context.Context, *Request, *Response) (err error)

type RestRequestHandler interface {
	GetConfig() ControllerConfig
	Create(*context.Context, *Request, *Response) (err error)
	Delete(*context.Context, *Request, *Response) (err error)
	Update(*context.Context, *Request, *Response) (err error)
	Get(*context.Context, *Request, *Response) (err error)
	List(*context.Context, *Request, *Response) (err error)
}

func (p *Gof) Router(path string, handler RequestHandler) {

}

func (p *Gof) Get(path string, handler RequestHandler) {

}

func (p *Gof) Post(path string, handler RequestHandler) {

}

func (p *Gof) Rest(path string, handler RestRequestHandler) {

}

func (p *Gof) loadConfig() {
	env := flag.String("e", "dev", "plz enter environment like dev test prod")
	p.SetEnv(*env)

	var path = fmt.Sprintf("./conf/%s/app.toml", *env)
	if _, err := toml.DecodeFile(path, &p.Config); err != nil {
		panic(err)
	}
}

func (p *Gof) initLogger() {
	file, err := os.OpenFile(p.Config.Logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Sprintf("%s:%s", err, p.Config.Logger.Path))
	}
	//defer file.Close()
	p.Logger.SetOutput(file)
}

func (p *Gof) SetLogger(logger log.Logger) {
	p.Logger = logger
}

type LoggerConf struct {
	Path string
}

type Conf struct {
	Logger LoggerConf
}

func (p *Gof) Start() {
	p.loadConfig()
	p.initLogger()

	p.Logger.Info("hello, i am monster!")
}
