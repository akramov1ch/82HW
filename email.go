package main

import "log"

func SendEmail(emailReq EmailRequest) error {
    log.Printf("Sending email to: %s, subject: %s, body: %s\n", emailReq.To, emailReq.Subject, emailReq.Body)
    return nil
}
