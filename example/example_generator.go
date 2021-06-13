package example

import (
	"github.com/go-specifier/specifier/lib"
)

// generated

func GetSpecification() (*lib.Specification, error) {

	// a creation of all dtos
	role := &Role{
		Entity: &lib.SpecParamBase{},
	}
	if err := role.Number.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.Number))

	if err := role.RootUserId.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.RootUserId))

	if err := role.Id.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.Id))

	user := &User{
		Entity: &lib.SpecParamBase{},
	}
	if err := user.Name.Init(); err != nil {
		return nil, err
	}
	user.With(lib.WithAttribute(&user.Name))

	if err := user.RoleId.Init(); err != nil {
		return nil, err
	}
	user.With(lib.WithAttribute(&user.RoleId))

	if err := user.Id.Init(); err != nil {
		return nil, err
	}
	user.With(lib.WithAttribute(&user.Id))

	// setup spec dtos
	data := &lib.Specification{}

	role.Setup(user)
	data.Add(role)

	user.Setup(role)
	data.Add(user)

	return data, nil
}
