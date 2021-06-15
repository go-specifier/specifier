package modelGenerator

import (
	"github.com/go-specifier/specifier/lib"
)

func WithType(name string) *OptionType {
	return &OptionType{
		typeName: name,
	}
}

func GetOptionType(in lib.SpecParam) (OptionType, bool) {
	for _, o := range in.Options() {
		if v, is := (o).(*OptionType); is {
			return *v, true
		}
	}
	return OptionType{}, false
}

var _ lib.SpecParamOption = (*OptionType)(nil)

type OptionType struct {
	typeName string
}

func (r OptionType) Is(option lib.SpecParamOption) {}

func (x OptionType) Multiple() bool {
	return false
}

func (x OptionType) String() string {
	return "type attribute: " + x.typeName
}

func (r OptionType) Name() string {
	return r.typeName
}
