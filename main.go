package main

import (
	"fmt"
)

func main() {
	data := GetData()
	for _, dto := range data.Dtos {
		fmt.Println("\ndto: " + dto.Name)
		for _, param := range dto.Params {
			fmt.Println("--- " + param.Name)
			for _, option := range param.Options {
				if ref, ok := option.(*Ref); ok {
					fmt.Printf("------ ref to:  %s.%s\n", ref.GetRef().Dto.Name, ref.GetRef().Name)

					for _, option := range ref.GetRef().Options {
						if t, ok := option.(*Type); ok {
							fmt.Println("--------- type:  " + t.Name())
						}
					}
				}
				if t, ok := option.(*Type); ok {
					fmt.Println("------ type:  " + t.Name())
				}
			}
		}
	}
}

// user code

func WithType(typeName string) *Type {
	return &Type{
		typeName: typeName,
	}
}

type Type struct {
	typeName string
}

func (r Type) Name() string {
	return r.typeName
}

func WithRef(ref *SpecParam) *Ref {
	return &Ref{
		ref: ref,
	}
}

type Ref struct {
	ref *SpecParam
}

func (r Ref) GetRef() *Attribute {
	return r.ref.outputAttribute
}

// user code

type User struct {
	Name   *SpecParam
	RoleId *SpecParam
}

func (u *User) Setup(r *Role) {
	u.Name = NewSpecParam(
		WithType("string"),
	)
	u.RoleId = NewSpecParam(
		WithRef(r.Id),
	)
}

type Role struct {
	Id *SpecParam
}

func (r *Role) Setup() {
	r.Id = NewSpecParam(
		WithType("int"),
	)
}

// generated

func GetData() Data {

	// a creation of all dtos
	role := &Role{}
	user := &User{}

	// setup spec dtos

	// FIXME - janhajek zalezi na poradi
	// obracene nefunguje, asi budou potreba takzvane narkovi dvojhvezdy
	role.Setup()
	user.Setup(role)

	// a creation of all dto output params
	userNameAttr := &Attribute{
		Name:    "Name",
		Options: []interface{}{},
	}
	user.Name.outputAttribute = userNameAttr

	userRoleIdAttr := &Attribute{
		Name:    "RoleId",
		Options: []interface{}{},
	}
	user.RoleId.outputAttribute = userRoleIdAttr

	roleIdAttribute := &Attribute{
		Name:    "Id",
		Options: []interface{}{},
	}
	role.Id.outputAttribute = roleIdAttribute

	//
	// FIXME - janhajek
	userNameAttr.Options = append(userNameAttr.Options, user.Name.options...)
	userRoleIdAttr.Options = append(userRoleIdAttr.Options, user.RoleId.options...)
	roleIdAttribute.Options = append(roleIdAttribute.Options, role.Id.options...)

	// output
	userDto := &Dto{
		Name:   "User",
		Params: []*Attribute{userNameAttr, userRoleIdAttr},
	}
	userNameAttr.Dto = userDto
	userRoleIdAttr.Dto = userDto

	roleDto := &Dto{
		Name:   "Role",
		Params: []*Attribute{roleIdAttribute},
	}
	roleIdAttribute.Dto = roleDto

	data := Data{
		Dtos: []*Dto{userDto, roleDto},
	}

	return data
}

// specifier

type SpecParam struct {
	options         []interface{}
	outputAttribute *Attribute
}

func NewSpecParam(o ...interface{}) *SpecParam {
	return &SpecParam{
		options: o,
	}
}

type Data struct {
	Dtos []*Dto
}

type Dto struct {
	Name   string
	Params []*Attribute
}

type Attribute struct {
	Name    string
	Options []interface{}
	Dto     *Dto
}
