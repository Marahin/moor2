package moor

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func AuthorizeByToken(w *http.ResponseWriter, r *http.Request) error {
	key, ok := r.URL.Query()["token"]

	if !ok || len(key[0]) < 1 || key[0] != AUTH_TOKEN {
		err := errors.New("URL parameter 'token' is invalid")

		log.WithFields(log.Fields{"token": key, "expected token": AUTH_TOKEN}).Error(err)
		http.Error(*w, err.Error(), http.StatusUnauthorized)

		return err
	}

	return nil
}
