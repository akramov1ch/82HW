package main

import (
    "encoding/json"
    "net/http"
)

type EmailRequest struct {
    To      string `json:"to"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
    var emailReq EmailRequest
    if err := json.NewDecoder(r.Body).Decode(&emailReq); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := SendEmail(emailReq); err != nil {
        http.Error(w, "Failed to send email", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Email sent successfully"))
}
