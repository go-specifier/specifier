package lib

import (
	"fmt"
)

type OptionEntity struct {
	SpecParam
}

var _ SpecParamOption = (*OptionEntity)(nil)

func (a OptionEntity) Is(_ SpecParamOption) {}

func (x OptionEntity) Multiple() bool {
	return false
}
func (x OptionEntity) String() string {
	return fmt.Sprintf("modelGenerator")
}

func WithEntity(param SpecParam) *OptionEntity {
	return &OptionEntity{
		SpecParam: param,
	}
}

func GetEntity(param SpecParam) (result *OptionEntity) {
	for _, o := range param.Options() {
		if v, is := (o).(*OptionEntity); is {
			return v
		}
	}
	return nil
}
