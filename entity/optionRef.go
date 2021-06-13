package entity

import (
	"github.com/go-specifier/specifier/lib"
)

var _ lib.SpecParamOption = (*optionRef)(nil)

type optionRef struct {
	to lib.SpecParam
}

func GetOptionRef(options ...lib.SpecParamOption) (optionRef, bool) {
	for _, o := range options {
		if v, is := (o).(*optionRef); is {
			return *v, true
		}
	}
	return optionRef{}, false
}

func (x optionRef) Is(lib.SpecParamOption) {}

func (x optionRef) Multiple() bool {
	return false
}

func (x optionRef) String() string {
	return "ref attribute"
}

func WithRef(ref lib.SpecParam) lib.SpecParamOption {
	return &optionRef{
		to: ref,
	}
}
