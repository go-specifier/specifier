package entity

import (
	"fmt"

	"github.com/go-specifier/specifier/lib"
)

func Generate(generator *lib.Specification) {
	for _, e := range generator.Entities {
		name, _ := GetOptionName(e.Options()...)
		fmt.Println("entity: ", name.name)
		for _, attr := range lib.GetAttributes(e.Options()...) {
			attrName, _ := GetOptionName(attr.Options()...)
			attrType, _ := GetOptionType(attr.Options()...)
			attrRef, hasAttrRef := GetOptionRef(attr.Options()...)
			fmt.Println("- attrName: ", attrName.name, attrType.Name())
			for _, o := range attr.Options() {
				fmt.Println("---- ", o.String())
			}
			if hasAttrRef {
				refName, _ := GetOptionName(attrRef.to.Options()...)
				fmt.Printf("- has ref to: %s\n", refName.name)
			}
		}
	}

}
