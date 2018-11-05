package captcha

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/mojocn/base64Captcha"
)

func GenerateCaptcha(cli *clientv3.Client, count, height, width int) (string, string, error) {
	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigCharacter{
		Height:             80,
		Width:              240,
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	if height > 0 {
		configD.Height = height
	}
	if width > 0 {
		configD.Width = width
	}

	//init etcd store
	base64Captcha.SetCustomStore(&EtcdStore{cli})

	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD, nil
}

func VerifyCaptcha(cli *clientv3.Client, id, captcha string) bool {
	base64Captcha.SetCustomStore(&EtcdStore{cli})
	return base64Captcha.VerifyCaptcha(id, captcha)
}
