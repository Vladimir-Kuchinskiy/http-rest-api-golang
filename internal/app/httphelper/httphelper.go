package httphelper

import (
	"encoding/json"
	"net/http"
)

type ctxKey int8

const (
	// CtxKeyUser ...
	CtxKeyUser ctxKey = iota
	// CtxKeyRequestID ...
	CtxKeyRequestID ctxKey = iota
)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, r *http.Request, code int, err error) {
	Respond(w, r, code, map[string]string{"error": err.Error()})
}

// Respond ...
func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
