package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	Username string `json:"username"`
	NickName string `json:"nickname"`
	Role     int    `json:"role"`
	UserID   uint   `json:"user_id"`
}
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var MySecret []byte
