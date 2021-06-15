package lib

import (
	"fmt"
)

type OptionAttribute struct {
	SpecParam
}

var _ SpecParamOption = (*OptionAttribute)(nil)

func (a OptionAttribute) Is(_ SpecParamOption) {}

func (x OptionAttribute) Multiple() bool {
	return true
}
func (x OptionAttribute) String() string {
	return fmt.Sprintf("attribute")
}

func WithAttribute(param SpecParam) *OptionAttribute {
	return &OptionAttribute{
		SpecParam: param,
	}
}

func GetAttributes(in SpecParam) (result []*OptionAttribute) {
	for _, o := range in.Options() {
		if v, is := (o).(*OptionAttribute); is {
			result = append(result, v)
		}
	}
	return
}
