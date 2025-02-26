package gomail

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendOptions struct {
	address    string
	name       string
	from       *string
	subject    *string
	templateId string
	items      map[string]string
}

func SendWithTemplate(options SendOptions) error {
	if config.Disabled {
		return nil
	}

	//Set Variables
	var from, subject string
	if options.from == nil {
		from = config.DefaultFrom
	} else {
		from = *options.from
	}
	if options.from == nil {
		subject = config.DefaultSubject
	} else {
		from = *options.subject
	}

	//Create Mail payload
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail(subject, from))
	m.SetTemplateID(options.templateId)

	//Adding Personalization
	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(options.name, options.address),
	}
	p.AddTos(tos...)
	for key, value := range options.items {
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
