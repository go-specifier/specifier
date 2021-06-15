package lib

type optionString string

var _ SpecParamOption = (*optionString)(nil)

func (a optionString) Is(_ SpecParamOption) {}

func (x optionString) Multiple() bool {
	return false
}
func (x optionString) String() string {
	return string(x)
}
