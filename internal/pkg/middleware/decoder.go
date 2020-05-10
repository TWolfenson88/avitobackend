package middleware

import (
	"avitocalls/internal/pkg/data"
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/network"
	"avitocalls/internal/pkg/settings"
	"io/ioutil"
	"net/http"
)

func DecodeBody(next settings.HandlerFunc) settings.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			network.Jsonify(w, forms.ErrorAnswer{
				Error:   err.Error(),
				Message: "Invalid Json",
			},  http.StatusNotAcceptable)
		}
		data.Body = body
		// _, err = io.Copy(ioutil.Discard, r.Body)
		next(w, r, ps)
	}
}
