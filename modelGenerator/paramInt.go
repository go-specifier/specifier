package modelGenerator

import (
	"github.com/go-specifier/specifier/lib"
)

type Int struct {
	lib.SpecParamBase
}

func (i *Int) Init() error {
	i.SpecParamBase.With(WithType("Int"))
	return i.SpecParamBase.Init()
}
