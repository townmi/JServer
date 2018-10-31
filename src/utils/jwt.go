package utils

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// JWT s
type JWT struct {
	SigningKey []byte
}

// Claims s
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var (
	signKey string = "test"
)

// NewJWT s
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey s
func GetSignKey() string {
	return signKey
}

// CreateToken s
func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// JWTAuth s
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, StandardFailMessage(tokenEmptyString(lang)))
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token, lang)
		if err != nil {
			c.JSON(http.StatusOK, StandardFailMessage(err.Error()))
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

// ParseToken s
func (j *JWT) ParseToken(tokenString string, lang string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(tokenValidationErrorString(lang))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(tokenValidationErrorExpiredString(lang))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(tokenValidationErrorNotValidYetString(lang))
			} else {
				return nil, errors.New(tokenHandleErrorString(lang))
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(tokenHandleErrorString(lang))
}

func tokenValidationErrorString(lang string) string {
	return GetLocalizer(lang, lang).MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TokenValidationError",
			Other: "That's not even a token",
		},
	})
}

func tokenValidationErrorExpiredString(lang string) string {
	return GetLocalizer(lang, lang).MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TokenValidationErrorExpired",
			Other: "Token is expired",
		},
	})
}

func tokenValidationErrorNotValidYetString(lang string) string {
	return GetLocalizer(lang, lang).MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TokenValidationErrorNotValidYet",
			Other: "Token not active yet",
		},
	})
}

func tokenHandleErrorString(lang string) string {
	return GetLocalizer(lang, lang).MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TokenHandleError",
			Other: "Couldn't handle this token",
		},
	})
}

func tokenEmptyString(lang string) string {
	return GetLocalizer(lang, lang).MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TokenEmptyError",
			Other: "请求未携带token，无权限访问",
		},
	})
}
