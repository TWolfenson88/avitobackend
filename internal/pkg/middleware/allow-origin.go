package middleware

import (
	"avitocalls/internal/pkg/settings"
	"net/http"
)

func SetAllowOrigin(next settings.HandlerFunc) settings.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}



