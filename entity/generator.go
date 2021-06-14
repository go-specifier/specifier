package entity

import (
	"fmt"

	"github.com/go-specifier/specifier/lib"
)

func Generate(generator *lib.Specification) {
	for _, e := range generator.Params {
		name, _ := GetOptionName(e)
		fmt.Println("entity: ", name.name)
		for _, attr := range lib.GetAttributes(e.Options()...) {
			attrName, _ := GetOptionName(attr)
			attrType, _ := GetOptionType(attr)
			//attrPackage := lib.GetPackage(attr)
			attrRef, hasAttrRef := GetOptionRef(attr)
			fmt.Println("- attrName: ", attrName.name, attrType.Name())
			for _, o := range attr.Options() {
				fmt.Println("---- ", o.String())
			}
			if hasAttrRef {
				refName, _ := GetOptionName(attrRef.to)
				refParent, _ := GetOptionName(lib.GetEntity(attrRef.to))
				fmt.Printf("- has ref to: %s.%s\n", refParent.name, refName.name)
			}
		}
	}

}
