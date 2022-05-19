package xhttp

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/unchain/json"
)

func Respond(w http.ResponseWriter, r *http.Request, v interface{}) {
	JSON(w, r, v)
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	bytes, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if status, ok := r.Context().Value(render.StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.Write(bytes)
}
