package entity

import (
	"github.com/go-specifier/specifier/lib"
)

var _ lib.SpecParamOption = (*OptionName)(nil)

type OptionName struct {
	name string
}

func GetOptionName(options ...lib.SpecParamOption) (OptionName, bool) {
	for _, o := range options {
		if v, is := (o).(*OptionName); is {
			return *v, true
		}
	}
	return OptionName{}, false
}

func (x OptionName) Is(lib.SpecParamOption) {}

func (x OptionName) Multiple() bool {
	return false
}

func (x OptionName) String() string {
	return "name attribute: " + x.name
}

func WithName(name string) lib.SpecParamOption {
	return &OptionName{
		name: name,
	}
}
