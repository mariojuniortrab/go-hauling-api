package handlers_protocols

import "net/http"

type Handle interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
