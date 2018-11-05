package user

// 验证码
type captchaRequest struct {
	NumCount int `json:"count"`
	Height   int `json:"height"`
	Width    int `json:"width"`
}

type captchaPic struct {
	Id      string `json:"id"`
	Captcha string `json:"captcha"`
}

type userlogin struct {
	User     string     `json:"user" description:"user name or email"`
	Password string     `json:"password" description:"user password"`
	Captcha  captchaPic `json:"captcha" description:"captcha"`
}

type userRegistry struct {
	User     string `json:"user" description:"user identifier for login"`
	Name     string `json:"name" description:"user name"`
	Password string `json:"password" description:"user password"`
	Email    string `json:"email" description:"email"`
	Profile  string `json:"profile" description:"profile"`
}
