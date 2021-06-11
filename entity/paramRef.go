package entity

import "github.com/go-specifier/specifier/lib"

type Ref struct {
	lib.SpecParamBase
}

func (i *Ref) Init() error {
	return i.SpecParamBase.Init()
}
