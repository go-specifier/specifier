package modelGenerator

import (
	"fmt"

	"github.com/go-specifier/specifier/lib"
)

func Generate(spec lib.SpecParam) {
	for _, e := range lib.GetEntities(spec) {
		name, _ := lib.GetOptionName(e)
		fmt.Println("modelGenerator: ", name.String())
		for _, attr := range lib.GetAttributes(e) {
			attrName, _ := lib.GetOptionName(attr)
			attrType, _ := GetOptionType(attr)
			attrPackage := lib.GetPackage(attr)
			attrRef, hasAttrRef := GetOptionRef(attr)
			fmt.Printf("- attrName: %s: %s in %s\n", attrName.String(), attrType.Name(), attrPackage.Name())
			for _, o := range attr.Options() {
				fmt.Println("---- ", o.String())
			}
			if hasAttrRef {
				refName, _ := lib.GetOptionName(attrRef.to)
				refPackage := lib.GetPackage(attrRef.to)
				refParent, _ := lib.GetOptionName(lib.GetEntity(attrRef.to))
				fmt.Printf("----  - has ref to: %s.%s from %s\n", refParent.String(), refName.String(), refPackage)
			}
		}
	}
}
