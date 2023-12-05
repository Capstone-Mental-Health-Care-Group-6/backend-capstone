package service

import (
	"FinalProject/configs"
	"FinalProject/features/users"
	"FinalProject/helper"
	"FinalProject/helper/enkrip"
	"errors"
	"math/rand"
	"strings"
	"time"
)

type UserService struct {
	d users.UserDataInterface
	j helper.JWTInterface
	c configs.ProgrammingConfig
	e enkrip.HashInterface
}

func New(data users.UserDataInterface, jwt helper.JWTInterface, config configs.ProgrammingConfig, enkrip enkrip.HashInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		j: jwt,
		c: config,
		e: enkrip,
	}
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	_, err := us.d.GetByEmail(newData.Email)
	if err == nil {
		return nil, errors.New("Email already registered by another user")
	}

	hashPassword, err := us.e.HashPassword(newData.Password)
	if err != nil {
		return nil, errors.New("Hash Password Error")
	}

	newData.Password = hashPassword

	result, err := us.d.Register(newData)

	if err != nil {
		return nil, errors.New("Failed to register")
	}
	return result, nil
}

func (us *UserService) Login(email, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("process failed")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func (us *UserService) GenerateJwt(email string) (*users.UserCredential, error) {
	result, err := us.d.GetByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("process failed")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func generateRandomCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

func (us *UserService) TokenResetVerify(code string) (*users.UserResetPass, error) {
	result, err := us.d.GetByCode(code)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("Token not found")
		}
		return nil, errors.New("Failed to verify token")
	}

	if result.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return result, nil
}

func (us *UserService) ForgetPasswordWeb(email string) error {

	user, err := us.d.GetByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("data not found")
		}
		return errors.New("process failed")
	}

	htmlBody := ""
	header := ""
	email = user.Email
	code := generateRandomCode(4)

	if user.Role == "Doctor" {

		url := us.c.BaseURLFE + "/reset-password?token_reset_password=" + code
		header = "Hi " + user.Name + ", Kami sudah mengirim link verifikasi untuk merubah kata sandi"
		htmlBody = `
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
							Jika Anda meminta pengaturan ulang kata sandi untuk @` + user.Name + `, konfirmasi atur ulang kata sandi dibawah untuk menyelesaikan proses permintaan ini, jika Anda tidak membuat permintaan ini, abaikan email ini.
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
	} else if user.Role == "Patient" {
		header = "Pemulihan Kata Sandi - Kode OTP Dikirimkan untuk Anda"
		htmlBody = `
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
							Halo, ` + user.Name + `
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
	}

	if err := us.d.InsertCode(email, code); err != nil {
		return err
	}

	ress := helper.SendEmail(email, header, htmlBody)

	if ress != nil {
		return ress
	}

	return nil
}

func (us *UserService) ResetPassword(code, email, password string) error {
	hashPassword, err := us.e.HashPassword(password)
	if err != nil {
		return errors.New("Hash Password Error")
	}
	password = hashPassword

	if err := us.d.ResetPassword(code, email, password); err != nil {
		return err
	}

	return nil
}
