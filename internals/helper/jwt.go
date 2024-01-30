package helper

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/masrayfa/configs"
	"github.com/masrayfa/internals/models/domain"
)

func ValidateToken(tokenString string) (user domain.UserRead, err error) {
	config := configs.GetConfig()

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		// return secret key
		return []byte(config.JWT.SecretKey), nil
	})

	if err != nil {
		return domain.UserRead{}, err
	}

	// validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.IdUser = int64(claims["id"].(float64))
		user.Email = claims["email"].(string)
		user.Username = claims["username"].(string)
		user.Status = claims["status"].(bool)
		user.IsAdmin = claims["is_admin"].(bool)
		return user, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return user, errors.New("Invalid token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return user, errors.New("Token expired")
	} else {
		return user, errors.New("Could not handle token")
	}
}
func ValidateUserCredentials(w http.ResponseWriter, r *http.Request) (domain.UserRead, error) {
	authorizationCookies, err := r.Cookie("authorization")
	if err != nil && err != http.ErrNoCookie {
		return domain.UserRead{}, err
	}

	authorizationCookiesValue := ""
	if authorizationCookies != nil {
		authorizationCookiesValue, err = url.QueryUnescape(authorizationCookies.Value)
		if err != nil {
			return domain.UserRead{}, err
		}
	}

	authorizationHeaders, haveAuthorizationHeader := r.Header["Authorization"]

	authorizationHeaderValue := ""
	if !haveAuthorizationHeader && authorizationCookiesValue == "" {
		return domain.UserRead{}, errors.New("Unauthorized")
	}

	if haveAuthorizationHeader {
		authorizationHeaderValue = authorizationHeaders[0]
	} else {
		authorizationHeaderValue = authorizationCookiesValue
	}

	authorizationSplit := strings.Fields(authorizationHeaderValue)
	if len(authorizationSplit) < 2 {
		return domain.UserRead{}, errors.New("Unauthorized")
	}

	authorizationType := authorizationSplit[0]
	if authorizationType != "Bearer" {
		return domain.UserRead{}, errors.New("Unauthorized")
	}

	authorizationToken := authorizationSplit[1]

	user, err := ValidateToken(authorizationToken)
	if err != nil {
		return domain.UserRead{}, err
	}

	return user, nil
}