package article

import "github.com/tomnattle/gof"
import "reflect"

func (p *Article) GetConfig() (config gof.ControllerConfig) {
	return gof.ControllerConfig{
		Denies: []reflect.Type{
			reflect.TypeOf(p.List),
			reflect.TypeOf(p.Delete),
		},
	}
}
