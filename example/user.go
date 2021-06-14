package example

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/lib"
)

type User struct {
	lib.SpecParam

	Id     entity.Int
	Name   entity.String
	RoleId entity.Ref
}

func (u *User) Setup(r *Role) {
	u.With(entity.WithName("ChangedUserName"))

	u.RoleId.With(
		entity.WithRef(&r.Id),
	)
}
