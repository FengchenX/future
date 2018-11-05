package explorer

// 资源管理器
type exDataSet struct {
	Id          string `json:"uuid,omitempty"`
	Name        string `json:"name"`
	UserId      string `json:"user_id"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type explorer struct {
	Path     string       `json:"path" description:"explorer path"`
	DataSets []*exDataSet `json:"newData, omitempty"`
}
