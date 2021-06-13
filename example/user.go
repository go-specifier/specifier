package example

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/lib"
)

type User struct {
	lib.Entity

	Id     entity.Int
	Name   entity.String
	RoleId entity.Ref
}

func (u User) Setup(r *Role) {
	u.With(
		entity.WithType("nic"),
	)

	u.Name.With(
		entity.WithType("blaType"),
	)
	u.RoleId.With(
		entity.WithRef(&r.Id),
	)
}
