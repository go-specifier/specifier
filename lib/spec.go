package lib

import (
	"reflect"
)

// specifier

type Entity interface {
	With(...SpecParamOption)
	Options() []SpecParamOption
}

type Generator interface {
	Generate()
}

type SpecParamOption interface {
	Is(option SpecParamOption)
	Multiple() bool
	String() string
}

type SpecParam interface {
	Init() error
	With(...SpecParamOption)
	Options() []SpecParamOption
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

func (i *SpecParamBase) With(options ...SpecParamOption) {
	for _, o := range options {
		if !o.Multiple() {
			var exists bool
			for index, savedO := range i.options {
				if reflect.TypeOf(savedO) == reflect.TypeOf(o) {
					exists = true
					i.options[index] = o
					break
				}
			}
			if exists {
				continue
			}
		}
		i.options = append(i.options, o)
	}
}

func (i *SpecParamBase) Options() []SpecParamOption {
	return i.options
}
