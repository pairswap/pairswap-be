package handler

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (h *SupportFormRequest) Send() error {
	from := mail.NewEmail("", os.Getenv("FROM_EMAIL"))
	subject := "User Form Contact"
	to := mail.NewEmail("", os.Getenv("TO_EMAIL"))
	plainTextContent := ""
	htmlContent := fmt.Sprintf("<div><span>Name: %s</span><br /><span>Email: %s</span><br /><span>Transaction: %s</span><br /><span>Comment: %s</span><br /></div>", h.Name, h.Email, h.TxURL, h.Comment)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)

	if err != nil {
		return err
	} else {
		return nil
	}
}
