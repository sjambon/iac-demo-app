package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iwert-m/terraform-azure/internal/config"
	"github.com/iwert-m/terraform-azure/internal/mailSender"
)

var Handler *handler

type handler struct {
	Configuration *config.Configuration
	MailSender    mailSender.IMailSender
}

func SetHandler(configuration *config.Configuration, mailSender mailSender.IMailSender) {
	if Handler == nil {
		Handler = &handler{
			Configuration: configuration,
			MailSender:    mailSender,
		}
	}
}

func (h *handler) HandleEventMailTrigger(w http.ResponseWriter, r *http.Request) {
	eventGridMessages := make([]event, 5)
	err := json.NewDecoder(r.Body).Decode(&eventGridMessages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, event := range eventGridMessages {
		if (event.Topic == nil) || (event.Subject == nil) || (event.EventType == nil) || (event.Data == nil) {
			log.Default().Printf("Invalid EventGrid message")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		mailAddress, ok := event.Data["email"]
		if !ok {
			log.Default().Printf("mailAddress is required")
			http.Error(w, "email is required", http.StatusBadRequest)
			return
		}

		err := h.MailSender.SendMail(mailAddress, "This mail was triggered by an EventGrid event")
		if err != nil {
			log.Default().Printf("Failed to send mail: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) HandlePostHttpMailTrigger(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Default().Printf("mailAddress is required")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mailAddress, ok := body["email"]
	if !ok {
		log.Default().Printf("mailAddress is required")
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	err = h.MailSender.SendMail(mailAddress, "This mail was triggered by an HTTP request")
	if err != nil {
		log.Default().Printf("Failed to send mail: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Mail sent!"))
}

func (h *handler) HandleGetHttpMailTrigger(w http.ResponseWriter, r *http.Request) {
	mailAddress := r.URL.Query().Get("email")
	if mailAddress == "" {
		w.Write([]byte("Send your email address in a query string with parameter <email> to receive a mail triggered by this app!"))
		return
	}

	err := h.MailSender.SendMail(mailAddress, "This mail was triggered by an HTTP request")
	if err != nil {
		log.Default().Printf("Failed to send mail: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Mail sent!"))
}
