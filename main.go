package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// handler untuk root /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// selain get diberi 404
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		// method GET akan mendapat OK
		response := map[string]string{
			"status": "OK",
		}
		// di-serialize
		jsonInBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// menulis response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("starting web server at http://localhost:%s/ \n", port)
	// berjalan di 8080
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
