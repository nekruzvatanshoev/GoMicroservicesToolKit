package errors

import "net/http"

func ResourceNotFound(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func InternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
}