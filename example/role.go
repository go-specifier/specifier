package example

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/lib"
)

type Role struct {
	lib.Entity

	Id         entity.Int
	Number     entity.Int
	RootUserId entity.Ref
}

func (r *Role) Setup(u *User) {
	r.With(
		entity.WithType("neco"),
	)

	r.Id.With(
		entity.WithType("int"),
	)
	r.RootUserId.With(
		entity.WithRef(&u.Id),
	)
}
