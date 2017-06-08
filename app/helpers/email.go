package helpers

import (
	"net/mail"
	"strings"

	"github.com/hjr265/postmark.go/postmark"
)

// EmailClient type
type EmailClient struct {
	Client *postmark.Client
}

// Email type
type Email struct {
	From    *mail.Address
	To      []*mail.Address
	Cc      []*mail.Address
	Bcc     []*mail.Address
	Subject string
	Body    string
}

// NewEmailClient func
func NewEmailClient(key string) *EmailClient {
	return &EmailClient{Client: &postmark.Client{ApiKey: key, Secure: true}}
}

// SendEmail func
func (client *EmailClient) SendEmail(email *Email) error {
	message := &postmark.Message{}
	message.From = email.From
	message.To = email.To
	message.Bcc = email.Bcc
	message.Subject = email.Subject
	message.HtmlBody = strings.NewReader(email.Body)

	_, err := client.Client.Send(message)
	return err
}
