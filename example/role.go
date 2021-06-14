package example

import (
	"github.com/go-specifier/specifier/entity"
	"github.com/go-specifier/specifier/lib"
)

type Role struct {
	lib.SpecParam

	Id           entity.Int
	Number       entity.Int
	RootUserId   entity.Ref
	ParentRoleId entity.Ref
	RootRoleId   entity.Ref
}

func (r *Role) Setup(u *User) {
	r.RootUserId.With(
		entity.WithRef(&u.Id),
	)
	r.RootRoleId.With(
		entity.WithRef(&r.Id),
	)
	r.ParentRoleId.With(
		entity.WithRef(&r.Id),
	)
}
