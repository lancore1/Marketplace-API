package model

type UserRole uint8

const (
	RoleAdmin UserRole = 0
	RoleUser  UserRole = 1
)

type UserInfo struct {
	ID   uint64
	Name string
	Role UserRole
}

func (u *UserInfo) GetRoleName() string {
	return UserRoles[u.Role]
}

var UserRoles = map[UserRole]string{
	RoleAdmin: "admin",
	RoleUser:  "user",
}
