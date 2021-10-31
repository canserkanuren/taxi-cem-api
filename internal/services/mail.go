package services

import (
	"cem-taxi-api/config"
	"cem-taxi-api/internal/models"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(body models.MailBody) string {
	defaultMail := mail.NewEmail("Taxi Ezanville / Mareil En France", config.GetConfig().Server.SendMail.DefaultMail)

	p := mail.NewPersonalization()
	p.AddTos(defaultMail)
	p.AddCCs(mail.NewEmail(fmt.Sprintf("%s %s", body.FirstName, body.LastName), body.From))

	m := new(mail.SGMailV3)
	m.SetFrom(defaultMail)
	m.Subject = body.Subject

	m.AddPersonalizations(p)
	var contents []*mail.Content

	contents = append(contents, mail.NewContent("text/html", body.Content))
	m.AddContent(contents...)

	client := sendgrid.NewSendClient(config.GetConfig().Server.SendMail.Key)
	response, _ := client.Send(m)

	statusCode := response.StatusCode

	if statusCode != 202 {
		log.Printf("Mail failed to be sent %v", response.Body)
		return response.Body
	} else {
		fmt.Printf("Mail is sent.. Response is: %v", response.Body)
		return ""
	}
}
