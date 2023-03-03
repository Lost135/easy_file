package tool

import (
	"easy_file/src/common"
	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("2324")

func GenerateJWT(claims common.Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Exp":        claims.Exp,
		"Authorized": claims.Authorized,
		"UserId":     claims.UserId,
		"Username":   claims.Username,
	})
	//token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
	//	"Exp":        claims.Exp,
	//	"Authorized": claims.Authorized,
	//	"UserId":     claims.UserId,
	//	"Username":   claims.Authorized,
	//})
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
