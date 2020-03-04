package moor

import (
	"github.com/google/jsonapi"

	"net/http"
)

func PrepareHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", jsonapi.MediaType)
	(*w).Header().Set("Access-Control-Allow-Origin", ALLOW_ORIGIN)
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, client, Cache-Control, X-Requested-(*w)ith, Access-Token, Uid")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
}
