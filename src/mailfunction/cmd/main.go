package main

import (
	"log"
	"net/http"

	"github.com/iwert-m/terraform-azure/internal/config"
	"github.com/iwert-m/terraform-azure/internal/handlers"
	"github.com/iwert-m/terraform-azure/internal/mailSender/mailjet"
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
	handlers.SetHandler(configuration, mailjet.NewMailJetManager(configuration))

	router.HandleFunc("/MailEventFunction", handlers.Handler.HandleEventMailTrigger)
	router.HandleFunc("POST /MailHttpFunction", handlers.Handler.HandlePostHttpMailTrigger)
	router.HandleFunc("GET /MailHttpFunction", handlers.Handler.HandleGetHttpMailTrigger)

	return router
}
