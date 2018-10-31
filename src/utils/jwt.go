package utils

// import "github.com/dgrijalva/jwt-go"

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	ID string `json:"id"`
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

func GetSignKey() string {
	return signKey
}
