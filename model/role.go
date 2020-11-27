package model

// Role enum
type Role int

const (
	// MEMBER role has access to just they infomation
	MEMBER Role = iota
	// ADMIN role has access to all infomation
	ADMIN
)

// ToString the role
func (r Role) ToString() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
