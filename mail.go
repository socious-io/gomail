package gomail

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendOptions struct {
	Address  string
	Name     string
	From     *string
	Subject  *string
	Template string
	Items    map[string]string
}

func SendWithTemplate(options SendOptions) error {
	if config.Disabled {
		return nil
	}

	//Set Variables
	var from, subject string
	if options.From == nil {
		from = config.DefaultFrom
	} else {
		from = *options.From
	}
	if options.From == nil {
		subject = config.DefaultSubject
	} else {
		from = *options.Subject
	}

	//Create Mail payload
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail(subject, from))
	m.SetTemplateID(options.Template)

	//Adding Personalization
	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(options.Name, options.Address),
	}
	p.AddTos(tos...)
	for key, value := range options.Items {
		p.SetDynamicTemplateData(key, value)
	}
	m.AddPersonalizations(p)

	//Setup the request
	request := sendgrid.GetRequest(config.ApiKey, "/v3/mail/send", config.Url)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)

	response, err := sendgrid.API(request)
	if err != nil {
		return err
	} else if strings.Split(strconv.Itoa(response.StatusCode), "")[0] != "2" {
		return errors.New(response.Body)
	}
	return nil
}

func GetTemplates() map[string]string {
	return config.Templates
}
