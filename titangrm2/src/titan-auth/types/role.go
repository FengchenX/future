package types

const (
	RoleAdmin      = "admin"
	RoleAnonymous  = "anonymous"
	RoleCommon     = "common"
	RoleGroupUser  = "group_user"
	RoleGroupAdmin = "group_admin"
)

type UserRole struct {
	User  User   `json:"user"`
	Roles []Role `json:"roles"`
}

// role
type Role struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
}
