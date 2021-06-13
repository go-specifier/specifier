package entity

import "github.com/go-specifier/specifier/lib"

type Ref struct {
	lib.SpecParamBase
}

func (i *Ref) Init() error {
	i.SpecParamBase.With(WithType("Ref"))
	return i.SpecParamBase.Init()
}
