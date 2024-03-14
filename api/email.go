package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendEmail(domain string, apiKey string, email string, mesg string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	sender := "mahdikaramy1117@gamil.com"
	subject := "Recommended songs"
	body := mesg
	recipient := email

	m := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return resp, err
}
