package dependencies

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/masrayfa/internals/models/domain"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator(validate *validator.Validate) *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}
// alternative:
// func NewValidator() *Validator {
// 	return &Validator{
// 		Validate: validator.New(),
// 	}
// }

func (v *Validator) validateStruct(req interface{}) error {
	err := v.Validate.Struct(req)
	return err
}

func (v *Validator) GetAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		potentialUser := request.Context().Value("currentUser")
		if potentialUser == nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, ok := potentialUser.(domain.UserRead)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("halo currentUser: ",user)
		// do something
		next.ServeHTTP(writer, request)
	})
}
