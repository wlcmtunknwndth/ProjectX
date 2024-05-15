package corsSkip

import "net/http"

func EnableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Max-Age", "86400")
	//w.WriteHeader(http.StatusOK)
}
