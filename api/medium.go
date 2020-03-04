package medium

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/jsonapi"
	log "github.com/sirupsen/logrus"

	"moor2/moor"
	. "moor2/pkg"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, client, Cache-Control, X-Requested-With, Access-Token, Uid")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
	// GET /api/spots
	if r.Method == http.MethodGet {
		log.Info("GET /api/medium")

		paramUrl := r.URL.Query()["url"][0]
		if StringInSlice(paramUrl, moor.IGNORE_ENDPOINTS) {
			fmt.Printf("  url ignored (%s)\n", paramUrl)
			return
		}

		parsedParamUrl, err := url.PathUnescape(paramUrl)
		if err != nil {
			fmt.Printf("Something went dang wrong, yo.")
			fmt.Printf("%s is unparsed paramUrl", paramUrl)
			fmt.Printf("%s is parsed paramUrl", parsedParamUrl)
			fmt.Print(err)
			_, _ = fmt.Fprint(w, "{ \"error\": \"PathUnescape failed\"}")
		}
		fmt.Printf("GET → parsed=%s\n", parsedParamUrl)
		_, _ = fmt.Fprintf(w, "%s", moor.Get(parsedParamUrl))

	} else {
		http.Error(w, errors.New("unsupported HTTP method").Error(), http.StatusInternalServerError)

		return
	}

}
