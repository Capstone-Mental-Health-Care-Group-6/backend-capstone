package helper

import "gopkg.in/gomail.v2"

func SendEmail(to, subject, body string) error {
	email := "irestu402@gmail.com"
	password := "irimwrkfzeuqlnsf"

	message := gomail.NewMessage()
	message.SetHeader("From", email)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	err := dialer.DialAndSend(message)
	if err != nil {
		return err
	}

	return nil
}
