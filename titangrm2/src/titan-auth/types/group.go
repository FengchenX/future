package types

// 组织
type Group struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentID    string `json:"parent_id,omitempty"`
	Path        string `json:"path,omitempty"`
	Layer       string `json:"layer,omitempty"`
	Level       int    `json:"level,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`

	Children []Group     `json:"children,omitempty"`
	Users    []GroupUser `json:"users,omitempty"`
}

// 组织里的用户
type GroupUser struct {
	GroupID  string `json:"group_id"`
	UserID   string `json:"user_id"`
	IsAdmin  bool   `json:"is_admin"`
	JoinTime string `json:"join_time"`

	Name       string `json:"name"`
	Type       string `json:"type"`
	CreateTime string `json:"create_time"`
	Email      string `json:"email"`
	Profile    string `json:"profile"`
	LastLogin  string `json:"last_login"`
	Status     string `json:"status"`
	Roles      []Role `json:"roles"`
}
