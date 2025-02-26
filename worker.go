package gomail

import (
	"fmt"
)

var EmailChannel = config.WorkerChannel

type EmailApproach string

const (
	EmailApproachTemplate EmailApproach = "TEMPLATE"
	EmailApproachDirect   EmailApproach = "DIRECT"
)

type EmailConfig struct {
	Approach    EmailApproach
	Destination string
	Title       string
	TemplateId  string
	Args        map[string]string
}

func SendEmail(emailConfig EmailConfig) {
	config.MessageQueue.SendJson(EmailChannel, emailConfig)
}

func EmailWorker(message interface{}) {
	emailConfig := new(EmailConfig)
	copy(message, emailConfig)

	var (
		destination = emailConfig.Destination
		title       = emailConfig.Title
		templateId  = emailConfig.TemplateId
		args        = emailConfig.Args
	)

	if emailConfig.Approach == EmailApproachTemplate {
		//Sending email with template
		err := SendWithTemplate(SendOptions{
			Address:  destination,
			Name:     title,
			Template: config.Templates[templateId],
			Items:    args,
		})
		if err != nil {
			fmt.Println("Coudn't Send Email, Error: ", err.Error())
		}
	}
}
