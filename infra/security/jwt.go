package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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

//GetEmailFromForgotPasswordToken return the emailToken jwt
func GetEmailFromForgotPasswordToken(tokenString string) (string, error) {
	if strings.TrimSpace(tokenString) == "" {
		return "", fmt.Errorf("token empty")
	}
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

//GenerateLoginToken return the token with user informations
func GenerateLoginToken(id int8, nick string, access int8) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     id,
		"nick":   nick,
		"access": access,
	})

	tokenString, err := token.SignedString(key)

	if err != nil {
		panic(err)
	}

	return tokenString
}

//GetClaimsFromToken return a array with id,nick and access from token
func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims, err
		}
	}
	return nil, err
}

//GetNickFromToken return a nick
func GetNickFromToken(tokenString string) (string, error) {
	claims, err := GetClaimsFromToken(tokenString)
	if err == nil {
		return claims["nick"].(string), nil
	}
	return "", nil
}

//GetNickFromRequest return a nick from a http request
func GetNickFromRequest(r *http.Request) string {
	cookie, err := r.Cookie("Authorization")
	if cookie != nil && err == nil {
		claims, err := GetClaimsFromToken(cookie.Value)

		if err == nil && claims["nick"] != nil && claims["nick"].(string) != "" {
			return claims["nick"].(string)
		}
	}
	return ""
}

//SetCookieToken set token on cookie with 7 days expiration
func SetCookieToken(w http.ResponseWriter, token string) {
	expire := time.Now().Add(7 * 24 * time.Hour) // Expires in 7 days
	cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", Expires: expire, MaxAge: 604800, HttpOnly: true, Secure: false}
	http.SetCookie(w, &cookie)
}

//SetExpiredCookie cookie with time expired to be deleted
func SetExpiredCookie(w http.ResponseWriter) {
	expire := time.Unix(0, 0)
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", Expires: expire, MaxAge: 0, HttpOnly: true, Secure: false}
	http.SetCookie(w, &cookie)
}

//IsValidCookie return true or false to valid cookie
func IsValidCookie(r *http.Request) bool {
	cookie, err := r.Cookie("Authorization")
	if cookie != nil && err == nil {
		claims, err := GetClaimsFromToken(cookie.Value)

		if err == nil && claims["nick"] != nil && claims["nick"].(string) != "" {
			return true
		}
	}
	return false
}
