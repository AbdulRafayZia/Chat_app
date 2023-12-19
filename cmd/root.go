package cmd

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/routes"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)


func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func Execute() {
	r := routes.Routes()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)
	port := 8080
	utils.Rafay()

	fmt.Printf("Server listening on :%d...\n", port)

	// Start the server
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
