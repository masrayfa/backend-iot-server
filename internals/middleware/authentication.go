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

func NewAuthenticationMiddleware(validator *dependencies.Validator) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		validator: validator,
	}
}

func (m *AuthenticationMiddleware) validateUserAndSetUserInHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		currentUser, err := helper.ValidateUserCredentials(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := request.Context()
		ctx = context.WithValue(ctx, "currentUser", currentUser)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}

func (m *AuthenticationMiddleware) ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
		if !ok {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(writer, request.WithContext(context.WithValue(ctx, "currentUser", currentUser)))
	})
}

func (m *AuthenticationMiddleware) ValidateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
		if !ok {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !currentUser.IsAdmin {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(writer, request.WithContext(context.WithValue(ctx, "currentUser", currentUser)))
	})
}

func (m *AuthenticationMiddleware) ValidateUserSameAsUrlIdOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id, err := m.validator.ParseIdFromUrlParameter(request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := request.Context()
		currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
		if !ok {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !currentUser.IsAdmin && currentUser.IdUser != id {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(writer, request.WithContext(context.WithValue(ctx, "currentUser", currentUser)))
	})
}