package types

const (
	User_UnActive  = "unactive"
	User_Active    = "active"
	User_Obsoleted = "obsoleted"

	User_Type_Guest      = "guest"
	User_Type_Individual = "individual"
	User_Type_Member     = "member"
	User_Type_Manager    = "manager"
	User_Type_Admin      = "admin"
)

// 用户
type User struct {
	Id       string `json:"id" description:"user id"`
	User     string `json:"user" description:"user identity"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	//Organization string `json:"organization,omitempty"`
	//Department   string `json:"deparrtment,omitempty"`
	Type       string `json:"type"`
	Profile    string `json:"profile"`
	CreateTime string `json:"create_time,omitempty"`
	LastLogin  string `json:"last_login"`
	Status     string `json:"status"`
	Session    string `json:"session,omitempty"`
}

// 用户
type UserList struct {
	Total int64   `json:"total"`
	Users []*User `json:"rows"`
}
