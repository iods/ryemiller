package server

import (
	"fmt"
	"github.com/iods/ryemiller.io/internal/env"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

func email(name string, email string, message string) {

	k := env.EnvVar("SENDGRID_KEY")

	from := mail.NewEmail("Your Website", "rye@drkstr.dev")
	subject := "From the website."
	to := mail.NewEmail("Example User", "millerrye17@gmail.com")
	content := mail.NewContent("text/plain", "A new email from " + name + "\n" + "Email: " + email + "\n" + "Message: " + message)
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(k, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}