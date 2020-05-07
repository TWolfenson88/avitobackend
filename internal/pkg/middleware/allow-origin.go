package middleware

import (
	"avitocalls/internal/pkg/settings"
	"net/http"
)

func SetAllowOrigin(next settings.HandlerFunc) settings.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next(w, r, ps)  //r.WithContext(ctx)
	}
}



