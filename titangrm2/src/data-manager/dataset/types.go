package dataset

type addDatasetRequest struct {
	Name        string `json:"name" description:"name of dataset"`
	Type        string `json:"type,omitempty" description:"type of dataset"`
	IsMarket    bool   `json:"market" description:"dataset of market"`
	Description string `json:"description,omitempty"`
}

type updateDatasetRequest struct {
	Name        string `json:"name" description:"name of dataset"`
	Description string `json:"description,omitempty"`
}

type addDatas struct {
	Datas []string `json:"datas" description:"data ids"`
}
