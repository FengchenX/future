package office

type login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type loginRes struct {
	Count      int `json:"count"`
	Status     int `json:"status"`
	StatusCode int `json:"statuscode"`
	Response   struct {
		Token   string `json:"token"`
		Expires string `json:"expires"`
		Sms     bool   `json:"sms"`
	} `json:"response"`
}

type uploadRes struct {
	Count      int `json:"count"`
	Status     int `json:"status"`
	StatusCode int `json:"statuscode"`
	Response   struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"response"`
}

type shareRes struct {
	Count      int    `json:"count"`
	Status     int    `json:"status"`
	StatusCode int    `json:"statuscode"`
	Response   string `json:"response"`
}
