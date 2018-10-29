package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var key string

// CreateMathCaptcha s
func CreateMathCaptcha(c *gin.Context) {
	var config = base64Captcha.ConfigCharacter{
		Height: 120,
		Width:  480,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexMedium,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexHigh,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         6,
	}
	secret, captchaInterfaceIntance := base64Captcha.GenerateCaptcha("", config)
	key = secret
	captchaInterfaceIntance.WriteTo(c.Writer)
}

func verfiyCaptcha(verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(key, verifyValue)
	return verifyResult
}
