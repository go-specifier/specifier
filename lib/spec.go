package lib

// specifier

type Entity interface {
	Option(interface{}) error
}

type Generator interface {
	Generate()
}

type SpecParamOption interface {
	X()
}

type SpecParam interface {
	Init() error
	With(...SpecParamOption)
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

func (i *SpecParamBase) Option(in interface{}) error {
	return nil
}
