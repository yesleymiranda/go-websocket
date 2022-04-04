package jwtoken

import "github.com/golang-jwt/jwt"

type Data struct {
	jwt.StandardClaims
	Username string `json:"username,omitempty"`
}
