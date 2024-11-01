package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var requestCount int
var requesttime time.Time

// var backendserver = []string{"localhost:8081", "localhost:8082", "localhost:8083"}

// func loadbalancerserver(backendserver []string) (balancer string) {
// 	for _, b := range backendserver {
// 		resp, err := http.Get(string(b))
// 		if err != nil || resp.StatusCode != http.StatusOK {
// 			log.Printf("Error connecting to backend server %s: %s", string(b), err)
// 			continue
// 		}
// 		defer resp.Body.Close()
// 	}
// 	return backendserver[0]
// }

func genericResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func logRequest(r *http.Request, status int) {
	log.Printf("request: %d --> method: %s --> path: %s --> Status-Code: %d", requestCount, r.Method, r.URL.Path, status)
}

func allAPIHandler(w http.ResponseWriter, r *http.Request) {
	if time.Since(requesttime).Seconds() > 1 {
		requesttime = time.Now()
		requestCount = 0
	}
	requestCount++
	if requestCount > 5 {
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		logRequest(r, http.StatusServiceUnavailable)
		return
	}

	switch r.URL.Path {
	case "/":
		http.Error(w, "stupid request", http.StatusBadRequest)
		logRequest(r, http.StatusBadRequest)
	case "/healthcheck":

		if _, err := os.Stat("/healthcheck"); err == nil {
			genericResponse(w, "Health Passed", http.StatusOK)
			logRequest(r, http.StatusOK)
		} else {
			genericResponse(w, "Health Failed", http.StatusInternalServerError)
			logRequest(r, http.StatusInternalServerError)
		}
	case "/api/v1/product":
		if _, err := os.Stat("/api/v1/product"); err == nil {
			genericResponse(w, "This is a Product API", http.StatusOK)
			logRequest(r, http.StatusOK)
		} else {
			genericResponse(w, "There is no Product API", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Invalid Path", http.StatusNotFound)
		logRequest(r, http.StatusNotFound)
	}

}
func main() {

	http.HandleFunc("GET /", allAPIHandler)
	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err1)
	}
}
