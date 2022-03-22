package handler

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (h *SupportFormRequest) Send() error {
	from := mail.NewEmail("Sender", os.Getenv("FROM_EMAIL"))
	subject := "Have a problem with transaction"
	to := mail.NewEmail("Receiver", os.Getenv("TO_EMAIL"))
	plainTextContent := "Example"
	htmlContent := fmt.Sprintf("<div><p>Name: %s<p><br /><p>Email: %s<p><br /><p>Transaction: %s<p><br /><p>Comment: %s<p><br /></div>", h.Name, h.Email, h.TxURL, h.Comment)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)

	if err != nil {
		return err
	} else {
		return nil
	}
}
