package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}
	log.Println(fmt.Sprintf("handler send message are %s", msg))
	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Send to " + requestPayload.To,
		Data:    msg.Data,
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}
