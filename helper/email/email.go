package email

import (
	"FinalProject/configs"
	"math/rand"

	"gopkg.in/gomail.v2"
)

type EmailInterface interface {
	SendEmail(to, subject, body string) error
	HTMLBody(role, name string) (string, string, string)
	HtmlBodyRegistDoctor(name string) (string, string)
}

type email struct {
	config configs.ProgrammingConfig
}

func New(c configs.ProgrammingConfig) EmailInterface {
	return &email{
		config: c,
	}
}

func (e *email) SendEmail(to, subject, body string) error {
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

func (e *email) generateRandomCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

func (e *email) HTMLBody(role, name string) (string, string, string) {
	code := e.generateRandomCode(4)
	var header, htmlBody string

	switch role {
	case "Doctor":
		header, htmlBody = e.htmlBodyDoctor(name, code)
		break
	case "Patient":
		header, htmlBody = e.htmlBodyPatient(name, code)
	}

	return header, htmlBody, code
}

func (e *email) htmlBodyDoctor(name, code string) (string, string) {
	url := e.config.BaseURLFE + "/reset-password/" + code
	header := "Hi " + name + ", Kami sudah mengirim link verifikasi untuk merubah kata sandi"
	htmlBody := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Kode Verifikasi</title>
		</head>
		<body style="margin: 0; padding: 0; box-sizing: border-box;">
			<table align="center" cellpadding="0" cellspacing="0" width="95%">
			<tr>
				<td align="center">
				<table align="center" cellpadding="0" cellspacing="0" width="600" style="border-spacing: 2px 5px;" bgcolor="#fff">
					<tr>
					<td style="padding: 5px 5px 5px 5px;">
						<a href="#" target="_blank">
						<img src="https://i.ibb.co/kgMjHSV/Logo.png" alt="Logo" style="width:200px; border:0; margin:0;"/>
						</a>
					</td>
					</tr>
					<tr>
					<td bgcolor="#fff">
						<table cellpadding="0" cellspacing="0" width="100%%">
						<tr>
							<td style="padding: 10px 0 10px 0; font-family: Nunito, sans-serif; font-size: 20px; font-weight: 900">
							Atur ulang kata sandi Anda? 
							</td>
						</tr>
						</table>
					</td>
					</tr>
					<tr>
					<td bgcolor="#fff">
						<table cellpadding="0" cellspacing="0" width="100%%">
						<tr>
							<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
							Jika Anda meminta pengaturan ulang kata sandi untuk @` + name + `, konfirmasi atur ulang kata sandi dibawah untuk menyelesaikan proses permintaan ini, jika Anda tidak membuat permintaan ini, abaikan email ini.
							</td>
						</tr>
						<tr>
							<td style="padding: 20px 0 20px 0; font-family: Nunito, sans-serif; font-size: 16px; text-align: center;">
							<a href="` + url + `" style="background-color: #0085FF; border: none; color: white; padding: 15px 40px; text-align: center; display: inline-block; font-family: Nunito, sans-serif; font-size: 18px; font-weight: bold; cursor: pointer; border-radius:8px; text-decoration:none">
								Atur Ulang Kata Sandi
							</a>
							</td>
						</tr>
						<tr>
							<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
							Jika Anda kesulitan mengklik tombol tersebut, copy dan paste URL di bawah ke dalam browser Anda:
							<a href="` + url + `">` + url + `</a>
							</td>
						</tr>
						<tr>
							<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
							Jika Anda tidak meminta pemulihan kata sandi ini, mohon abaikan pesan ini untuk menjaga keamanan akun Anda. 
							</td>
						</tr>
						</table>
					</td>
					</tr>
				</table>
				</td>
			</tr>
			</table>
		</body>
		</html>
	`

	return header, htmlBody
}

func (e *email) htmlBodyPatient(name, code string) (string, string) {
	header := "Pemulihan Kata Sandi - Kode OTP Dikirimkan untuk Anda"
	htmlBody := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Kode Verifikasi</title>
		</head>
		<body style="margin: 0; padding: 0; box-sizing: border-box;">
			<table align="center" cellpadding="0" cellspacing="0" width="95%">
			<tr>
				<td align="center">
				<table align="center" cellpadding="0" cellspacing="0" width="600" style="border-spacing: 2px 5px;" bgcolor="#fff">
					<tr>
					<td style="padding: 5px 5px 5px 5px;">
						<a href="#" target="_blank">
						<img src="https://i.ibb.co/kgMjHSV/Logo.png" alt="Logo" style="width:200px; border:0; margin:0;"/>
						</a>
					</td>
					</tr>
					<tr>
					<td bgcolor="#fff">
						<table cellpadding="0" cellspacing="0" width="100%%">
						<tr>
							<td style="padding: 10px 0 10px 0; font-family: Nunito, sans-serif; font-size: 20px; font-weight: 900">
							Halo, ` + name + `
							</td>
						</tr>
						</table>
					</td>
					</tr>
					<tr>
					<td bgcolor="#fff">
						<table cellpadding="0" cellspacing="0" width="100%%">
						<tr>
							<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
							Kami melihat bahwa Anda mengalami kesulitan untuk mengakses akun Anda. Jangan khawatir, kami di sini untuk membantu Anda! Kami telah mengirimkan kode OTP ke alamat email terkait dengan akun Anda.
							</td>
						</tr>
						<tr>
							<td style="padding: 20px 0 20px 0; font-family: Nunito, sans-serif; font-size: 16px; text-align: center;">
							<button style="background-color: #0085FF; border: none; color: white; padding: 15px 30px; text-align: center; display: inline-block; font-family: Nunito, sans-serif; font-size: 20px; font-weight: bold; cursor: pointer; border-radius:8px; letter-spacing: 10px;">
								` + code + `
							</button>
							</td>
						</tr>
						<tr>
							<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
							Silakan gunakan kode ini untuk mengatur ulang kata sandi Anda dengan mudah. Pastikan untuk segera mengganti kata sandi setelah berhasil masuk kembali ke akun Anda.
							<p></p>
							</td>
						</tr>
						<tr>
							<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
							Jika Anda tidak meminta pemulihan kata sandi ini, mohon abaikan pesan ini untuk menjaga keamanan akun Anda. 
							</td>
						</tr>
						</table>
					</td>
					</tr>
				</table>
				</td>
			</tr>
			</table>
		</body>
		</html>
		`

	return header, htmlBody
}

func (e *email) HtmlBodyRegistDoctor(name string) (string, string) {
	header := "Selamat " + name + ", pengajuan konselor Anda sudah kami terima!"
	htmlBody := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Pengajuan Konselor</title>
	</head>
	<body style="margin: 0; padding: 0; box-sizing: border-box;">
		<table align="center" cellpadding="0" cellspacing="0" width="95%">
		<tr>
			<td align="center">
			<table align="center" cellpadding="0" cellspacing="0" width="600" style="border-spacing: 2px 5px;" bgcolor="#fff">
				<tr>
				<td style="padding: 5px 5px 5px 5px;">
					<a href="#" target="_blank">
					<img src="https://i.ibb.co/kgMjHSV/Logo.png" alt="Logo" style="width:200px; border:0; margin:0;"/>
					</a>
				</td>
				</tr>
				<tr>
				<td bgcolor="#fff">
					<table cellpadding="0" cellspacing="0" width="100%%">
					<tr>
						<td style="padding: 10px 0 10px 0; font-family: Nunito, sans-serif; font-size: 18px; font-weight: 900">
						Halo, ` + name + `
						</td>
					</tr>
					</table>
				</td>
				</tr>
				<tr>
				<td bgcolor="#fff">
					<table cellpadding="0" cellspacing="0" width="100%%">
					<tr>
						<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
						Selamat! Kami dengan senang hati ingin memberitahu Anda bahwa pengajuan Anda sebagai Konselor di EmpathiCare telah berhasil diterima. Kami sangat berterima kasih atas ketertarikan Anda untuk bergabung dengan kami.
			<p></p>
						</td>
					</tr>
					<tr>
						<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
						Terima kasih atas kontribusi Anda dalam meningkatkan pelayanan kesehatan kami.
			 <p></p>
						</td>
					</tr>
		  <tr>
						<td style="padding: 10px 0 10px 0; font-family: Nunito, sans-serif; font-size: 16px; font-weight: 900">
						Salam Sehat,
						</td>
					</tr>
		   <tr>
						<td style="font-family: Nunito, sans-serif; font-size: 16px; font-weight: 900">
						Team EmpathiCare
						</td>
					</tr>
					</table>
				</td>
				</tr>
			</table>
			</td>
		</tr>
		</table>
	</body>
	</html>`

	return header, htmlBody
}
