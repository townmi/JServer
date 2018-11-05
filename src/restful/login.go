package restful

import (
	"net/http"
	"time"

	models "../models"
	sqls "../sqls"
	utils "../utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type loginResponse struct {
	Token    string `json:"token"`
	UserInfo user   `json:"userInfo"`
}

type user struct {
	ID       string `json:"id"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Level    int    `json:"level"`
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
}

// AccountLogin s
func AccountLogin(c *gin.Context) {
	account := map[string]string{
		"email":    "",
		"password": "",
		// "captcha":  "",
	}
	err := utils.ValidateRequestForm(c, account)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	// if !verfiyCaptcha(account["captcha"]) {
	// 	c.JSON(http.StatusOK, utils.StandardErrorMessage("captcha"))
	// 	return
	// }

	u, err := sqls.ValidateUser(account)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}
	res, err := generateWithTokenResponse(u)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}
	// cookie := &http.Cookie{
	// 	Name:     "id_token",
	// 	Value:    res.Token,
	// 	Path:     "/",
	// 	Domain:   c.ConfGet("server.domain"),
	// 	HttpOnly: true,
	// 	MaxAge:   10,
	// }
	// http.SetCookie(c.Writer, cookie)

	c.SetCookie("id_token", res.Token, 30, "/", "goubaa.com", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   res,
		"code":   1,
	})
}

func generateWithTokenResponse(sqlUser *models.User) (loginResponse, error) {
	j := &utils.JWT{
		SigningKey: []byte(utils.SignKey),
	}
	claims := utils.Claims{
		ID:       sqlUser.ID,
		UserName: sqlUser.UserName,
		Mobile:   sqlUser.Mobile,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    utils.SignKey,                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return loginResponse{}, err
	}

	_u := user{
		ID:       sqlUser.ID,
		Mobile:   sqlUser.Mobile,
		Email:    sqlUser.Email,
		Level:    sqlUser.Level,
		UserName: sqlUser.UserName,
		Avatar:   sqlUser.Avatar,
	}
	return loginResponse{
		Token:    token,
		UserInfo: _u,
	}, nil
}
