package mailjet

import (
	"strings"

	"github.com/iwert-m/terraform-azure/internal/config"
	"github.com/mailjet/mailjet-apiv3-go"
)

type MailJetManager struct {
	configuration *config.Configuration
	mailClient    *mailjet.Client
}

func NewMailJetManager(configuration *config.Configuration) *MailJetManager {
	mailClient := mailjet.NewMailjetClient(
		configuration.MailServiceConfiguration.ApiKey,
		configuration.MailServiceConfiguration.SecretKey)

	return &MailJetManager{
		configuration: configuration,
		mailClient:    mailClient,
	}
}

func (m *MailJetManager) SendMail(mailAddress string, subject string) error {
	receiverName := strings.Split(mailAddress, ".")[0]
	if receiverName == mailAddress {
		receiverName = strings.Split(mailAddress, "@")[0]
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: m.configuration.MailServiceConfiguration.SenderMail,
				Name:  "iac-demo-mailer",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: mailAddress,
					Name:  receiverName,
				},
			},
			Subject:    subject,
			TemplateID: 6290674,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	_, err := m.mailClient.SendMailV31(&messages)
	if err != nil {
		return err
	}

	return nil
}
