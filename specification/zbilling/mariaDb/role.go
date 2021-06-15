package mariaDb

import (
	"github.com/go-specifier/specifier/lib"
	"github.com/go-specifier/specifier/modelGenerator"
)

type Role struct {
	lib.SpecParam

	Id           modelGenerator.Int
	Number       modelGenerator.Int
	RootUserId   modelGenerator.Ref
	ParentRoleId modelGenerator.Ref
	RootRoleId   modelGenerator.Ref
}

func (r *Role) Setup(u *User) {
	r.RootUserId.With(
		modelGenerator.WithRef(&u.Id),
	)
	r.RootRoleId.With(
		modelGenerator.WithRef(&r.Id),
	)
	r.ParentRoleId.With(
		modelGenerator.WithRef(&r.Id),
	)
}
