package lib

import (
	"fmt"
)

type OptionEntities struct {
	SpecParam
}

var _ SpecParamOption = (*OptionEntities)(nil)

func (a OptionEntities) Is(_ SpecParamOption) {}

func (x OptionEntities) Multiple() bool {
	return true
}
func (x OptionEntities) String() string {
	return fmt.Sprintf("entities")
}

func WithEntities(param SpecParam) *OptionEntities {
	return &OptionEntities{
		SpecParam: param,
	}
}

func GetEntities(param SpecParam) (result []*OptionEntities) {
	for _, o := range param.Options() {
		if v, is := (o).(*OptionEntities); is {
			result = append(result, v)
		}
	}
	return
}
