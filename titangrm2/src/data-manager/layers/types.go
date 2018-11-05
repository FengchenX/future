package layers

type addLayerReq struct {
	Name        string `json:"layer_name"`
	Style       string `json:"style_id"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default" description:"false"`
}
