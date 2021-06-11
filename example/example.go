package example

import (
	"github.com/go-specifier/specifier/entity"
)

type User struct {
	Id     entity.Int
	Name   entity.String
	RoleId entity.Ref
}

func (u User) Setup(r *Role) {
	u.Name.With(
		entity.WithType("blaType"),
	)
	u.RoleId.With(
		entity.WithRef(&r.Id),
	)
}

type Role struct {
	Id         entity.Int
	Number     entity.Int
	RootUserId entity.Ref
}

func (r *Role) Setup(u *User) {
	r.Id.With(
		entity.WithType("int"),
	)
	r.RootUserId.With(
		entity.WithRef(&u.Id),
	)
}
