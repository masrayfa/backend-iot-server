package middleware

import (
	"context"
	"net/http"

	"github.com/masrayfa/internals/dependencies"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
)

type AuthenticationMiddleware struct {
	validator *dependencies.Validator
}

func NewAuthenticationMiddleware(validator *dependencies.Validator) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		validator: validator,
	}
}

type Result struct {
	User domain.UserRead
	Err error
}

func (m *AuthenticationMiddleware) validateUserAndSetUserInHeader(next http.Handler, writer http.ResponseWriter, request *http.Request) Result {
	currentUser, err := helper.ValidateUserCredentials(writer, request)
	return Result {
		User: currentUser,
		Err: err,
	}
}

func (m *AuthenticationMiddleware) ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		result := m.validateUserAndSetUserInHeader(next, writer, request)
		if result.Err != nil {
			http.Error(writer, result.Err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := request.Context()
		ctx = context.WithValue(ctx, "currentUser", result.User)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}

func (m *AuthenticationMiddleware) ValidateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		result := m.validateUserAndSetUserInHeader(next, writer, request)
		if result.Err != nil {
			http.Error(writer, result.Err.Error(), http.StatusUnauthorized)
			return
		}

		if !result.User.IsAdmin {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := request.Context()
		ctx = context.WithValue(ctx, "currentUser", result.User)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}

func (m *AuthenticationMiddleware) ValidateUserSameAsUrlIdOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id, err := m.validator.ParseIdFromUrlParameter(request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		result := m.validateUserAndSetUserInHeader(next, writer, request)
		if result.Err != nil {
			http.Error(writer, result.Err.Error(), http.StatusUnauthorized)
			return
		}

		if !result.User.IsAdmin && result.User.IdUser != id {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := request.Context()
		ctx = context.WithValue(ctx, "currentUser", result.User)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}