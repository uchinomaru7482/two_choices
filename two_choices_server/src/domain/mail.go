package domain

import (
	"context"
	"fmt"
	"net/http"
	"two_choices/pkg/sendgrid"

	"github.com/pkg/errors"
	sendgridgo "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// メール
type Mail struct {
	FromName         string
	FromEmail        string
	ToName           string
	ToEmail          string
	Subject          string
	PlainTextContent string
	HTMLContent      string
}

// FromName - 送信者名
const FromName string = "管理者"

// FromEmail - 送信元アドレス
const FromEmail string = "test@example.com"

// Send - メール送信
func (s *Mail) Send(ctx context.Context, config *sendgrid.Config) error {
	from := mail.NewEmail(s.FromName, s.FromEmail)
	to := mail.NewEmail(s.ToName, s.ToEmail)
	message := mail.NewSingleEmail(from, s.Subject, to, s.PlainTextContent, s.HTMLContent)

	client := sendgridgo.NewSendClient(config.APIKey)
	res, err := client.Send(message)
	if err != nil {
		return err
	}
	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusAccepted) {
		return errors.New(fmt.Sprintf("mail send error:%v", res))
	}
	return nil
}
