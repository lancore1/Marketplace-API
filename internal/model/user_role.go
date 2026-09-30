package model

import (
	"fmt"
	"strings"
)

type RoleType string

type UserRole struct {
	RoleID int32    `db:"role_id" json:"role_id"`
	Role   RoleType `db:"role" json:"role_type"`
}

const (
	RoleAdmin RoleType = "admin"
	RoleUser  RoleType = "user"
)

func (r RoleType) Valid() bool {
	str := RoleType(strings.ToLower(strings.TrimSpace(string(r))))
	fmt.Println(r)
	switch str {
	case RoleAdmin, RoleUser:
		return true
	default:
		return false
	}

}
