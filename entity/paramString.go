package entity

import (
	"github.com/go-specifier/specifier/lib"
)

type String struct {
	lib.SpecParamBase
}

func (i *String) Init() error {
	return i.SpecParamBase.Init()
}
