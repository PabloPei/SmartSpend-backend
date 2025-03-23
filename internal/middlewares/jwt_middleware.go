package middlewares

import (
	"context"
	"net/http"

	"github.com/PabloPei/SmartSpend-backend/internal/auth"
	"github.com/PabloPei/SmartSpend-backend/internal/errors"
	"github.com/PabloPei/SmartSpend-backend/internal/models"
	"github.com/PabloPei/SmartSpend-backend/utils"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, userService models.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := utils.GetTokenFromRequest(r)

		claims, err := auth.ValidateJWT(tokenString, false)
		if err != nil {
			utils.WriteError(w, http.StatusForbidden, errors.ErrJWTInvalidToken)
			return
		}

		email := claims["email"].(string)

		u, err := userService.GetUserPublicByEmail(email)
		if err != nil {
			utils.WriteError(w, http.StatusForbidden, errors.ErrUserNotFound)
			return
		}

		// Add the user to the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, models.UserKey, u.UserId)
		r = r.WithContext(ctx)

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

// TODO agregar validacion del access token que este vencido pero pertenezca al usuario
func WithRefreshTokenAuth(handlerFunc http.HandlerFunc, userService models.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := utils.GetTokenFromRequest(r)

		claims, err := auth.ValidateJWT(tokenString, true)
		if err != nil {
			utils.WriteError(w, http.StatusForbidden, errors.ErrJWTInvalidToken)
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			utils.WriteError(w, http.StatusForbidden, errors.ErrJWTInvalidToken)
			return
		}

		u, err := userService.GetUserPublicByEmail(email)
		if err != nil {
			utils.WriteError(w, http.StatusForbidden, errors.ErrUserNotFound)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, models.UserKey, u.UserId)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
