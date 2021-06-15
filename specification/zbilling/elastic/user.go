package elastic

import (
	"github.com/go-specifier/specifier/lib"
	"github.com/go-specifier/specifier/modelGenerator"
)

type User struct {
	lib.SpecParam

	Id     modelGenerator.Int
	Name   modelGenerator.String
	RoleId modelGenerator.Ref
}

func (u *User) Setup(r *Role) {
	u.With(lib.WithName("ChangedUserName"))

	u.RoleId.With(
		modelGenerator.WithRef(&r.Id),
	)
}
