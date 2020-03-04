package medium

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"

	"moor2/moor"
	. "moor2/pkg"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	moor.PrepareHeaders(&w)
	if err := moor.AuthorizeByToken(&w, r); err != nil {
		return
	}

	// GET /api/medium?url=A&token=B
	if r.Method == http.MethodGet {
		log.WithFields(log.Fields{"r": r}).Info("GET /api/medium")

		paramUrl := r.URL.Query()["url"][0]
		if StringInSlice(paramUrl, moor.IGNORE_ENDPOINTS) {
			fmt.Printf("  url ignored (%s)\n", paramUrl)
			return
		}

		parsedParamUrl, err := url.PathUnescape(paramUrl)
		if err != nil {
			log.WithFields(log.Fields{
				"unparsed paramUrl": paramUrl,
				"parsed paramUrl":   parsedParamUrl,
			}).Error(err)
			_, _ = fmt.Fprint(w, "{ \"error\": \"PathUnescape failed\"}", http.StatusInternalServerError)
		}

		_, _ = fmt.Fprintf(w, "%s", moor.Get(parsedParamUrl))
		log.WithFields(log.Fields{"parsed paramUrl": parsedParamUrl}).Info("GET /api/medium completed")

		return
	} else {
		http.Error(w, errors.New("unsupported HTTP method").Error(), http.StatusInternalServerError)

		return
	}
}
