package security

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var key = []byte("THENOOBSRPG!!!")

//GenerateForgotPasswordToken return the emailToken jwt
func GenerateForgotPasswordToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	tokenString, err := token.SignedString(key)

	if err != nil {
		panic(err)
	}

	return tokenString
}

//GenerateLoginToken return the token with user informations
func GenerateLoginToken(id int8, nick string, email string, access int8) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     id,
		"nick":   nick,
		"email":  email,
		"access": access,
	})

	tokenString, err := token.SignedString(key)

	if err != nil {
		panic(err)
	}

	return tokenString
}

//GetEmailFromForgotPasswordToken return the emailToken jwt
func GetEmailFromForgotPasswordToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	email := ""
	if token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email = claims["email"].(string)
			fmt.Println(email)
		}
	}
	return email, err
}
