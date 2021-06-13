package entity

import (
	"fmt"

	"github.com/go-specifier/specifier/lib"
)

func Generate(generator *lib.Specification) {
	for _, e := range generator.Entities {
		name := optionName{}
		e.Option(&name)

		fmt.Println("entity: ", name.name)
		//for _, param := range dto.Params {
		//	fmt.Println("--- " + param.Name)
		//	for _, option := range param.Options {
		//		if ref, ok := option.(*Ref); ok {
		//			fmt.Printf("------ ref to:  %s.%s\n", ref.GetRef().Dto.Name, ref.GetRef().Name)
		//
		//			for _, option := range ref.GetRef().Options {
		//				if t, ok := option.(*Type); ok {
		//					fmt.Println("--------- type:  " + t.Name())
		//				}
		//			}
		//		}
		//		if t, ok := option.(*Type); ok {
		//			fmt.Println("------ type:  " + t.Name())
		//		}
		//	}
		//}
	}

}
