package modelGenerator

import (
	"github.com/go-specifier/specifier/lib"
)

type String struct {
	lib.SpecParamBase
}

func (i *String) Init() error {
	i.SpecParamBase.With(WithType("String"))
	return i.SpecParamBase.Init()
}
