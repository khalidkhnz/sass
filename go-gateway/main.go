package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/khalidkhnz/sass/go-gateway/config"
)

func main() {
	// Initialize environment variables
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
	path := r.URL.Path

	switch {
	case strings.HasPrefix(path, "/go-blog"):
		proxyRequest(w, r, config.GetBlogUrl("http://localhost:8081"),"/go-blog")
	case strings.HasPrefix(path, "/go-ecom"):
		proxyRequest(w, r, config.GetEcomUrl("http://localhost:8082"),"/go-ecom")
	case strings.HasPrefix(path, "/go-sass"):
		proxyRequest(w, r, config.GetSassUrl("http://localhost:8083"),"/go-sass")
	default:
		sendJSONError(w, http.StatusNotFound, "Route not found")
	}
}
func proxyRequest(w http.ResponseWriter, r *http.Request, backendURL string, havePrefix string) {
	// Parse the backend URL
	backend, err := url.Parse(backendURL)
	if err != nil {
		log.Printf("Invalid backend URL: %v\n", err)
		sendJSONError(w, http.StatusInternalServerError, "Invalid backend URL")
		return
	}

	// Remove service prefix from path
	path := r.URL.Path
	path = strings.TrimPrefix(path, havePrefix)

	r.URL.Path = path

	// Resolve the full backend URL, including query parameters
	targetURL := backend.ResolveReference(r.URL)

	// Create a new request to the backend service
	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		log.Printf("Failed to create proxy request: %v\n", err)
		sendJSONError(w, http.StatusInternalServerError, "Failed to create proxy request")
		return
	}

	// Copy headers from the original request
	proxyReq.Header = r.Header

	// Use an HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Printf("Failed to contact backend service: %v\n", err)
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
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("Failed to write response body: %v\n", err)
	}
}

func sendJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Prepare JSON response
	response := map[string]interface{}{
		"success": false,
		"message": message,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Failed to send JSON error response: %v\n", err)
	}
}
