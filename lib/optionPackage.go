package lib

type OptionPackage = optionString

var _ SpecParamOption = (*OptionPackage)(nil)

func (x OptionPackage) Name() string {
	return x.String()
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
