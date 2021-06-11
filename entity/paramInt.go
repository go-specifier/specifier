package entity

import (
	"github.com/go-specifier/specifier/lib"
)

type Int struct {
	lib.SpecParamBase
}

func (i *Int) Init() error {
	return i.SpecParamBase.Init()
}
