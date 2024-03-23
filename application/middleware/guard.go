package guard

import (
	"antrein/bc-queue/model/config"
	"antrein/bc-queue/model/dto"
	"antrein/bc-queue/model/entity"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type GuardContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type AuthGuardContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Claims         entity.CustomClaim
}

func (g *GuardContext) ReturnError(status int, message string) error {
	g.ResponseWriter.WriteHeader(status)
	return json.NewEncoder(g.ResponseWriter).Encode(dto.NoBodyDTOResponseWrapper{
		Status:  status,
		Message: message,
	})
}

func (g *GuardContext) ReturnSuccess(body interface{}) error {
	g.ResponseWriter.WriteHeader(http.StatusOK)
	return json.NewEncoder(g.ResponseWriter).Encode(dto.DefaultDTOResponseWrapper{
		Status:  http.StatusOK,
		Message: "ok",
		Body:    body,
	})
}

func (g *AuthGuardContext) ReturnError(status int, message string) error {
	g.ResponseWriter.WriteHeader(status)
	return json.NewEncoder(g.ResponseWriter).Encode(dto.NoBodyDTOResponseWrapper{
		Status:  status,
		Message: message,
	})
}

func (g *AuthGuardContext) ReturnSuccess(body interface{}) error {
	g.ResponseWriter.WriteHeader(http.StatusOK)
	return json.NewEncoder(g.ResponseWriter).Encode(dto.DefaultDTOResponseWrapper{
		Status:  http.StatusOK,
		Message: "ok",
		Body:    body,
	})
}

func DefaultGuard(handlerFunc func(g *GuardContext) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guardCtx := GuardContext{
			ResponseWriter: w,
			Request:        r,
		}
		if err := handlerFunc(&guardCtx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func AuthGuard(cfg *config.Config, handlerFunc func(g *AuthGuardContext) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization") // or another method, depending on your token location
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Secrets.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authGuardCtx := AuthGuardContext{
			ResponseWriter: w,
			Request:        r,
			Claims:         entity.CustomClaim{},
		}

		if err := handlerFunc(&authGuardCtx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
