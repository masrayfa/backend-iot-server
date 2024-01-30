package dependencies

import (
	"context"
	"log"
	"net/http"
	"strconv"

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

func (v *Validator) ParseIdFromUrlParameter(r *http.Request) (int64, error) {
	potentialId := r.Context().Value("id")
	if potentialId == nil {
		param := r.URL.Query().Get("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			log.Println("error parsing id: ", err)
			return 0, err
		}

		if param == "" {
			log.Println("error parsing id: ", err)
		}

		r = r.WithContext(context.WithValue(r.Context(), "id", id))

		return id, nil
	} else {
		id, ok := potentialId.(int64)
		if !ok {
			log.Println("error parsing id: ", ok)
			return 0, nil
		}

		return id, nil
	}
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
