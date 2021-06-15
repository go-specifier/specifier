package elastic

import (
	"github.com/go-specifier/specifier/lib"
)

type Spec struct {
	lib.SpecParam
	User *User
	Role *Role
}

func (s *Spec) Setup() error {

	*s = Spec{
		SpecParam: &lib.SpecParamBase{},
		Role: &Role{
			SpecParam: &lib.SpecParamBase{},
		},
		User: &User{
			SpecParam: &lib.SpecParamBase{},
		},
	}

	// a creation of all dtos
	role := s.Role
	{
		role.With(lib.WithName("Role"))
		role.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
		role.With(lib.WithEntity(s))
		{
			role.With(lib.WithAttribute(&role.Number))
			if err := role.Number.Init(); err != nil {
				return err
			}
			role.Number.With(lib.WithEntity(role))
			role.Number.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			role.Number.With(lib.WithName("Number"))
		}
		{
			role.With(lib.WithAttribute(&role.RootUserId))
			if err := role.RootUserId.Init(); err != nil {
				return err
			}
			role.RootUserId.With(lib.WithEntity(role))
			role.RootUserId.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			role.RootUserId.With(lib.WithName("RootUserId"))
		}
		{
			role.With(lib.WithAttribute(&role.RootRoleId))
			if err := role.RootRoleId.Init(); err != nil {
				return err
			}
			role.RootRoleId.With(lib.WithEntity(role))
			role.RootRoleId.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			role.RootRoleId.With(lib.WithName("RootRoleId"))
		}
		{
			role.With(lib.WithAttribute(&role.ParentRoleId))
			if err := role.ParentRoleId.Init(); err != nil {
				return err
			}
			role.ParentRoleId.With(lib.WithEntity(role))
			role.ParentRoleId.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			role.ParentRoleId.With(lib.WithName("RootRoleId"))
			role.With(lib.WithAttribute(&role.ParentRoleId))
		}
		{
			role.With(lib.WithAttribute(&role.Id))
			if err := role.Id.Init(); err != nil {
				return err
			}
			role.Id.With(lib.WithEntity(role))
			role.Id.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			role.Id.With(lib.WithName("Id"))
		}
	}
	user := s.User

	user.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
	user.With(lib.WithEntity(s))
	{
		user.With(lib.WithName("User"))
		user.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))

		{
			user.With(lib.WithAttribute(&user.Id))
			if err := user.Id.Init(); err != nil {
				return err
			}
			user.Id.With(lib.WithEntity(user))
			user.Id.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			user.Id.With(lib.WithName("Id"))
		}
		{
			if err := user.Name.Init(); err != nil {
				return err
			}
			user.With(lib.WithAttribute(&user.Name))
			user.Name.With(lib.WithEntity(user))
			user.Name.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			user.Name.With(lib.WithName("Name"))
		}
		{
			if err := user.RoleId.Init(); err != nil {
				return err
			}
			user.With(lib.WithAttribute(&user.RoleId))
			user.RoleId.With(lib.WithEntity(user))
			user.RoleId.With(lib.WithPackage("github.com/go-specifier/specifier/zbilling/elastic"))
			user.RoleId.With(lib.WithName("RoleId"))
		}
	}

	role.Setup(user)
	s.With(lib.WithEntities(role))

	user.Setup(role)
	s.With(lib.WithEntities(user))

	return nil
}
