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
	{
		role.With(entity.WithName("Role"))
		role.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
		{
			role.With(lib.WithAttribute(&role.Number))
			if err := role.Number.Init(); err != nil {
				return nil, err
			}
			role.Number.With(lib.WithEntity(role))
			role.Number.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			role.Number.With(entity.WithName("Number"))
		}
		{
			role.With(lib.WithAttribute(&role.RootUserId))
			if err := role.RootUserId.Init(); err != nil {
				return nil, err
			}
			role.RootUserId.With(lib.WithEntity(role))
			role.RootUserId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			role.RootUserId.With(entity.WithName("RootUserId"))
		}
		{
			role.With(lib.WithAttribute(&role.RootRoleId))
			if err := role.RootRoleId.Init(); err != nil {
				return nil, err
			}
			role.RootRoleId.With(lib.WithEntity(role))
			role.RootRoleId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			role.RootRoleId.With(entity.WithName("RootRoleId"))
		}
		{
			role.With(lib.WithAttribute(&role.ParentRoleId))
			if err := role.ParentRoleId.Init(); err != nil {
				return nil, err
			}
			role.ParentRoleId.With(lib.WithEntity(role))
			role.ParentRoleId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			role.ParentRoleId.With(entity.WithName("RootRoleId"))
			role.With(lib.WithAttribute(&role.ParentRoleId))
		}
		{
			role.With(lib.WithAttribute(&role.Id))
			if err := role.Id.Init(); err != nil {
				return nil, err
			}
			role.Id.With(lib.WithEntity(role))
			role.Id.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			role.Id.With(entity.WithName("Id"))
		}
	}
	user := &User{
		SpecParam: &lib.SpecParamBase{},
	}
	{
		user.With(entity.WithName("User"))
		user.With(lib.WithPackage("github.com/go-specifier/specifier/example"))

		{
			user.With(lib.WithAttribute(&user.Id))
			if err := user.Id.Init(); err != nil {
				return nil, err
			}
			user.Id.With(lib.WithEntity(user))
			user.Id.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			user.Id.With(entity.WithName("Id"))
		}
		{
			if err := user.Name.Init(); err != nil {
				return nil, err
			}
			user.With(lib.WithAttribute(&user.Name))
			user.Name.With(lib.WithEntity(user))
			user.Name.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			user.Name.With(entity.WithName("Name"))
		}
		{
			if err := user.RoleId.Init(); err != nil {
				return nil, err
			}
			user.With(lib.WithAttribute(&user.RoleId))
			user.RoleId.With(lib.WithEntity(user))
			user.RoleId.With(lib.WithPackage("github.com/go-specifier/specifier/example"))
			user.RoleId.With(entity.WithName("RoleId"))
		}
	}
	// setup spec dtos
	data := &lib.Specification{}

	role.Setup(user)
	data.Add(role)

	user.Setup(role)
	data.Add(user)

	return data, nil
}
