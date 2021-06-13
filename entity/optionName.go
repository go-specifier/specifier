package entity

import (
	"github.com/go-specifier/specifier/lib"
)

var _ lib.SpecParamOption = (*optionName)(nil)

type optionName struct {
	name string
}

func (x optionName) Is(option lib.SpecParamOption) bool {
	return false
}

func WithName(name string) lib.SpecParamOption {
	return &optionName{
		name: name,
	}
}
