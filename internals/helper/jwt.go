package helper

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/masrayfa/configs"
	"github.com/masrayfa/internals/models/domain"
)

// func SignUserToken(user domain.UserRead) (string, error) {
// 	config := configs.GetConfig()

// 	log.Println("user dari fungsi sign user token: ", user)

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id_user": user.IdUser,
// 		"username": user.Username,
// 		"email": user.Email,
// 		"status": user.Status,
// 		"isAdmin": user.IsAdmin,
// 		"iat": time.Now().Unix(),
// 	})

// 	tokenString, err := token.SignedString([]byte(config.JWT.SecretKey))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func SignUserToken(user domain.UserRead) (string, error) {
	config := configs.GetConfig()
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"idUser":   user.IdUser,
		"email":    user.Email,
		"username": user.Username,
		"status":   user.Status,
		"isAdmin":  user.IsAdmin,
		"iat":      time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.JWT.SecretKey))

	return tokenString, err
}


func ValidateToken(tokenString string) (user domain.UserRead, err error) {
	config := configs.GetConfig()
	// Parse and validate the token. KeyFunc will be used to validate the token by returning the secret key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm used in the token is the same we used to sign it
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT.SecretKey), nil
	})
	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.IdUser = int64(claims["idUser"].(float64))
		user.Email = claims["email"].(string)
		user.Username = claims["username"].(string)
		user.Status = claims["status"].(bool)
		user.IsAdmin = claims["isAdmin"].(bool)
		return user, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return user, errors.New("invalid token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return user, errors.New("token expired")
	} else {
		return user, errors.New("could not handle token")
	}
}


// func ValidateToken(tokenString string) (user domain.UserRead, err error) {
// 	config := configs.GetConfig()

// 	// parse token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// validate signing method
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid token")
// 		}

// 		// return secret key
// 		return []byte(config.JWT.SecretKey), nil
// 	})

// 	if err != nil {
// 		return domain.UserRead{}, err
// 	}

// 	log.Println("token dari fungsi validate token: ", token)

// 	// validate claims
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		user.IdUser = int64(claims["id_user"].(float64))
// 		user.Email = claims["email"].(string)
// 		user.Username = claims["username"].(string)
// 		user.Status = claims["status"].(bool)
// 		user.IsAdmin = claims["isAdmin"].(bool)
// 		return user, nil
// 	} else if errors.Is(err, jwt.ErrTokenMalformed) {
// 		return user, errors.New("invalid token")
// 	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
// 		return user, errors.New("token expired")
// 	} else {
// 		return user, errors.New("could not handle token")
// 	}
// }

func ValidateUserCredentials(w http.ResponseWriter, r *http.Request) (domain.UserRead, error) {
	authorizationCookies, err := r.Cookie("authorization")
	if err != nil && err != http.ErrNoCookie {
		return domain.UserRead{}, err
	}

	if authorizationCookies != nil {
		log.Println("authorizationCookies name: ", authorizationCookies.Name)
	} else {
		log.Println("authorizationCookies is nil")
	}

	authorizationCookiesValue := ""
	if authorizationCookies != nil {
		authorizationCookiesValue, err = url.QueryUnescape(authorizationCookies.Value)
		if err != nil {
			return domain.UserRead{}, err
		}
	} else {
		return domain.UserRead{}, errors.New("Unauthorized")
	}

	log.Println("authorizationCookiesValue: ", authorizationCookiesValue)

	authorizationHeaders, haveAuthorizationHeader := r.Header["Authorization"]

	log.Println("authorizationHeaders: ", authorizationHeaders)

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

	log.Println("authorizationSplit[0]: ", authorizationSplit[0])

	authorizationType := authorizationSplit[0]
	if authorizationType != "Bearer" {
		return domain.UserRead{}, errors.New("Unauthorized")
	}

	log.Println("authorizationSplit[1]: ", authorizationSplit[1])

	authorizationToken := authorizationSplit[1]

	user, err := ValidateToken(authorizationToken)
	if err != nil {
		return domain.UserRead{}, err
	}

	log.Println("user dari middleware validate user credentials: ", user)

	return user, nil


	// headerAuth := r.Header.Get("Authorization")
	// log.Println("headerAuth: ", headerAuth)
}