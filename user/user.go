package user

type UserRole int

const (
	RoleAdmin UserRole = 0
	RoleUser  UserRole = 1
)

type UserInfo struct {
	ID   int
	Name string
	Role UserRole
}

func (u *UserInfo) GetRoleName() string {
	return RoleMap[u.Role]
}

const RoleMap = map[UserRole]string{
	RoleAdmin: "admin",
	RoleUser:  "user",
}
