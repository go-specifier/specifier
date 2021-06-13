package entity

import (
	"github.com/go-specifier/specifier/lib"
)

func WithType(name string) *Type {
	return &Type{
		typeName: name,
	}
}

var _ lib.SpecParamOption = (*Type)(nil)

type Type struct {
	typeName string
}

func (r Type) Is(option lib.SpecParamOption) bool {
	return false
}

func (r Type) Name() string {
	return r.typeName
}
