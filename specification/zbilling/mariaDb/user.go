package mariaDb

import (
	"github.com/go-specifier/specifier/lib"
	"github.com/go-specifier/specifier/modelGenerator"
	"github.com/go-specifier/specifier/specification/zbilling/elastic"
)

type User struct {
	lib.SpecParam

	Id     modelGenerator.Int
	Name   modelGenerator.String
	RoleId modelGenerator.Ref
}

func (u *User) Setup(r *Role, user *elastic.User) {
	u.With(lib.WithName("ChangedUserName"))

	u.RoleId.With(
		modelGenerator.WithRef(&user.RoleId),
	)
}
