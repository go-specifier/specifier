package example

import (
	"github.com/go-specifier/specifier/lib"
)

// generated

func GetSpecification() (*lib.Specification, error) {

	// a creation of all dtos
	role := &Role{}
	var roleAttributes []lib.Entity
	if err := role.Number.Init(); err != nil {
		return nil, err
	}
	roleAttributes = append(roleAttributes, &role.Number)

	if err := role.RootUserId.Init(); err != nil {
		return nil, err
	}
	roleAttributes = append(roleAttributes, &role.RootUserId)

	if err := role.Id.Init(); err != nil {
		return nil, err
	}
	roleAttributes = append(roleAttributes, &role.Id)

	user := &User{}
	var userAttributes []lib.Entity
	if err := user.Name.Init(); err != nil {
		return nil, err
	}
	userAttributes = append(userAttributes, &user.Name)
	if err := user.RoleId.Init(); err != nil {
		return nil, err
	}
	userAttributes = append(userAttributes, &user.RoleId)
	if err := user.Id.Init(); err != nil {
		return nil, err
	}
	userAttributes = append(userAttributes, &user.Id)

	// setup spec dtos
	data := &lib.Specification{}

	role.Setup(user)
	data.Add(role)

	user.Setup(role)
	data.Add(user)

	return data, nil
}
