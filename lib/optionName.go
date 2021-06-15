package lib

var _ SpecParamOption = (*OptionName)(nil)

type OptionName struct {
	name string
}

func GetOptionName(in SpecParam) (OptionName, bool) {
	for _, o := range in.Options() {
		if v, is := (o).(*OptionName); is {
			return *v, true
		}
	}
	return OptionName{}, false
}

func (x OptionName) Is(SpecParamOption) {}

func (x OptionName) Multiple() bool {
	return false
}

func (x OptionName) String() string {
	return "name attribute: " + x.name
}

func WithName(name string) SpecParamOption {
	return &OptionName{
		name: name,
	}
}
