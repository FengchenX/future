package styles

type addStyleReq struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Sld         string `json:"sld"`
	Type        string `json:"type" description:"point/line/polygon"`
}
