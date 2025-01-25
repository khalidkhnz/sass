package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/khalidkhnz/sass/go-gateway/config"
)

func main() {
	config.InitEnv()

	fmt.Println("API GATEWAY IS UP AND RUNNING")

	// Create a new ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", proxyHandler)

	// Create handler chain with middleware
	handler := LoggingMiddleware(mux)

	log.Println("API Gateway is running on port 8080...")
	log.Fatal(http.ListenAndServe(config.GetPort(), handler))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Define the backend services
	routes := map[string]string{
		config.GetBlogPostFix("/go-blog"): config.GetBlogUrl("http://localhost:8081"),
		config.GetEcomPostFix("/go-ecom"): config.GetEcomUrl("http://localhost:8082"),
		config.GetSassPostFix("go-sass"): config.GetSassUrl("http://localhost:8083"),
	}

	// Match the incoming request path
	for route, backendURL := range routes {
		if r.URL.Path == route {
			proxyRequest(w, r, backendURL)
			return
		}
	}

	// Default response for unmatched paths
	sendJSONError(w, http.StatusNotFound, "Route not found")
}

func proxyRequest(w http.ResponseWriter, r *http.Request, backendURL string) {
	// Parse the backend URL
	backend, err := url.Parse(backendURL)
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Invalid backend URL")
		return
	}

	// Create a new request to the backend service
	proxyReq, err := http.NewRequest(r.Method, backend.String(), r.Body)
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Failed to create proxy request")
		return
	}

	// Copy headers from the original request
	proxyReq.Header = r.Header

	// Forward the request to the backend
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		sendJSONError(w, http.StatusBadGateway, "Failed to contact backend service")
		return
	}
	defer resp.Body.Close()

	// Copy response headers and body back to the original client
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func sendJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Prepare JSON response
	response := map[string]interface{}{
		"success": false,
		"message": message,
	}
	json.NewEncoder(w).Encode(response)
}
