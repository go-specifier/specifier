package main

import (
	"fmt"

	"github.com/go-specifier/specifier/modelGenerator"
	"github.com/go-specifier/specifier/specification/zbilling/elastic"
	"github.com/go-specifier/specifier/specification/zbilling/mariaDb"
)

func main() {
	elasticSpec := &elastic.Spec{}
	elasticSpec.Setup()

	mariaDbSpec := &mariaDb.Spec{}
	mariaDbSpec.Setup(elasticSpec)

	fmt.Println("MariaDb SPEC")
	modelGenerator.Generate(mariaDbSpec)
	fmt.Println("---------------------------------------")
	fmt.Println("Elastic SPEC")
	modelGenerator.Generate(elasticSpec)
	fmt.Println("---------------------------------------")
}
