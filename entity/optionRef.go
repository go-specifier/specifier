package entity

import (
	"github.com/go-specifier/specifier/lib"
)

var _ lib.SpecParamOption = (*optionRef)(nil)

type optionRef struct {
	to lib.SpecParam
}

func (x optionRef) Is(option lib.SpecParamOption) bool {
	return false
}

func WithRef(ref lib.SpecParam) lib.SpecParamOption {
	return &optionRef{
		to: ref,
	}
}
