package action

import "net/http"

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("health check OK!"))
	w.WriteHeader(http.StatusOK)
}
