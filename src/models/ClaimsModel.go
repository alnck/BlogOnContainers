package models

import (
	. "blog-on-containers/constants"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Username string `json:"username,omitempty"`
	Roles    []int  `json:"roles,omitempty"`
	jwt.StandardClaims
}

const ip = "127.0.0.1"

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf(MESSAGE_TOKEN_INVALID)
}

func (claims JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
