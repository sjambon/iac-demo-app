package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/iwert-m/terraform-azure/internal/config"
)

func main() {
	configuration := config.New()

	srv := &http.Server{
		Addr:    ":" + configuration.Port,
		Handler: routes(configuration),
	}

	log.Default().Printf("Listening on %s", configuration.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func routes(configuration *config.Configuration) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /RelayHttpFunction", func(w http.ResponseWriter, r *http.Request) {
		mailAddress := r.URL.Query().Get("email")
		if mailAddress == "" {
			w.Write([]byte("Send your email address in a query string with parameter <email> to receive a mail triggered by this app!"))
			return
		}

		requestPayload := map[string]string{"email": mailAddress}
		jsonBody, err := json.Marshal(requestPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = http.Post("https://"+configuration.RelayTargetUrl+"/MailHttpFunction", "application/json", bytes.NewReader(jsonBody))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte("Request relayed!"))
	})

	return router
}
