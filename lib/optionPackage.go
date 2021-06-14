package lib

import (
	"fmt"
)

type OptionPackage string

var _ SpecParamOption = (*OptionPackage)(nil)

func (a OptionPackage) Is(_ SpecParamOption) {}

func (x OptionPackage) Multiple() bool {
	return false
}
func (x OptionPackage) String() string {
	return fmt.Sprintf("package: %s", string(x))
}

func WithPackage(name string) *OptionPackage {
	result := OptionPackage(name)
	return &result
}

func GetPackage(param SpecParam) (result *OptionPackage) {
	for _, o := range param.Options() {
		if v, is := (o).(*OptionPackage); is {
			return v
		}
	}
	return nil
}
