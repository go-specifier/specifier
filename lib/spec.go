package lib

// specifier

type Entity interface {
	With(...SpecParamOption)
	Option(SpecParamOption) (SpecParamOption, error)
	Options([]interface{}) error
}

type Generator interface {
	Generate()
}

type SpecParamOption interface {
	Is(option SpecParamOption) bool
}

type SpecParam interface {
	Init() error
	With(...SpecParamOption)
}

type SpecEntity struct {
	options []SpecParamOption
}

type SpecParamBase struct {
	options []SpecParamOption
}

func (i *SpecParamBase) Init() error {
	return nil
}

func (i *SpecParamBase) With(option ...SpecParamOption) {
	i.options = append(i.options, option...)
}

func (i *SpecParamBase) Option(in SpecParamOption) (SpecParamOption, error) {
	for _, o := range i.options {
		if o.Is(in) {
			return o, nil
		}
	}
	return nil, nil
}

func (i *SpecParamBase) Options(in []interface{}) error {
	return nil
}
