package example

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/lib"
)

// generated

func GetSpecification() (*lib.Specification, error) {

	// a creation of all dtos
	role := &Role{
		SpecParam: &lib.SpecParamBase{},
	}
	role.With(entity.WithName("Role"))
	role.With(lib.WithPackage("github.com/go-specifier/specifier/example"))

	role.Number.With(lib.WithEntity(role))
	role.Number.With(entity.WithName("Number"))
	role.Number.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	role.With(lib.WithAttribute(&role.Number))

	role.RootUserId.With(lib.WithEntity(role))
	role.RootUserId.With(entity.WithName("RootUserId"))
	role.RootUserId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	if err := role.RootUserId.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.RootUserId))

	role.RootRoleId.With(lib.WithEntity(role))
	role.RootRoleId.With(entity.WithName("RootRoleId"))
	role.RootRoleId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	if err := role.RootRoleId.Init(); err != nil {
		return nil, err
	}

	role.With(lib.WithAttribute(&role.ParentRoleId))
	role.ParentRoleId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	role.ParentRoleId.With(lib.WithEntity(role))
	role.ParentRoleId.With(entity.WithName("RootRoleId"))
	if err := role.ParentRoleId.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.ParentRoleId))

	role.Id.With(lib.WithEntity(role))
	role.Id.With(entity.WithName("Id"))
	role.Id.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	if err := role.Id.Init(); err != nil {
		return nil, err
	}
	role.With(lib.WithAttribute(&role.Id))

	user := &User{
		SpecParam: &lib.SpecParamBase{},
	}
	user.With(entity.WithName("User"))
	user.With(lib.WithPackage("github.com/go-specifier/specifier/example"))

	user.Id.With(lib.WithEntity(user))
	user.Id.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	user.Id.With(entity.WithName("Id"))
	if err := user.Name.Init(); err != nil {
		return nil, err
	}
	user.Name.With(entity.WithName("Name"))
	user.With(lib.WithAttribute(&user.Name))

	if err := user.RoleId.Init(); err != nil {
		return nil, err
	}
	user.RoleId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
	user.RoleId.With(lib.WithEntity(user))
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
