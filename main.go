package main

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/example"
)

func main() {
	spec, err := example.GetSpecification()
	if err != nil {
		panic(err)
	}
	entity.Generate(spec)
}
