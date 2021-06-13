package example

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/lib"
)

// generated

func GetSpecification() (*lib.Specification, error) {

	// a creation of all dtos
	role := &Role{
		Entity: &lib.SpecParamBase{},
	}
	role.With(entity.WithName("Role"))

	role.Number.With(entity.WithName("Number"))
	role.With(lib.WithAttribute(&role.Number))

	role.RootUserId.With(entity.WithName("RootUserId"))
	if err := role.RootUserId.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.RootUserId))

	role.Id.With(entity.WithName("Id"))
	if err := role.Id.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.Id))

	user := &User{
		Entity: &lib.SpecParamBase{},
	}
	user.With(entity.WithName("User"))

	user.Id.With(entity.WithName("Id"))
	if err := user.Name.Init(); err != nil {
		return nil, err
	}
	user.Name.With(entity.WithName("Name"))
	user.With(lib.WithAttribute(&user.Name))

	if err := user.RoleId.Init(); err != nil {
		return nil, err
	}
	user.RoleId.With(entity.WithName("RoleId"))
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
