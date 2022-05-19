package xhttp

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/unchain/json"
)

func Decode(r *http.Request, v interface{}) error {
	return DecodeJSON(r.Body, v)
}

func DecodeFormValue(r *http.Request, key string, v interface{}) error {
	err := json.Unmarshal([]byte(r.FormValue(key)), v)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func DecodeJSON(r io.Reader, v interface{}) error {
	defer io.Copy(ioutil.Discard, r)
	return json.NewDecoder(r).Decode(v)
}
