package lib

type AttributeOption struct {
	param SpecParam
}

var _ SpecParamOption = (*AttributeOption)(nil)

func (a AttributeOption) Is(option SpecParamOption) bool {
	return false
}

func WithAttribute(param SpecParam) *AttributeOption {
	return &AttributeOption{
		param: param,
	}
}
