package mailsender

import (
	"fmt"
	"net/smtp"
	"shop/models"
)

type MailAPI interface {
	SendOrderNotification(order *models.Order) error
}

type mailAPI struct {
	server, sender, password string
}

func NewMailAPI(server, sender, password string) (*mailAPI, error) {
	return &mailAPI{
		server:   server,
		sender:   sender,
		password: password,
	}, nil
}

func (m *mailAPI) SendOrderNotification(order *models.Order) error {
	auth := smtp.PlainAuth("", m.sender, m.password, m.server)
	message := fmt.Sprintf("To: <%s>\nFrom: <%s>\nSubject: Ваш заказ принят\n\nВаш заказ с номером %v принят в работу.\nВ ближайшее время с Вами свяжется менеджер.", order.Email, m.sender, order.ID)

	err := smtp.SendMail(m.server+":587", auth, m.sender, []string{order.Email}, []byte(message))
	return err
}
