package service

import (
	"FinalProject/features/users"
	"FinalProject/features/users/mocks"
	helper "FinalProject/helper/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)
	newUser := users.User{
		Name:     "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "password",
		Role:     "Admin",
		Status:   "Active",
	}

	t.Run("Success Register", func(t *testing.T) {
		data.On("GetByEmail", newUser.Email).Return(&newUser, errors.New("")).Once()
		enkrip.On("HashPassword", newUser.Password).Return(newUser.Password, nil).Once()
		data.On("Register", newUser).Return(&newUser, nil).Once()

		result, err := service.Register(newUser)

		data.AssertExpectations(t)
		enkrip.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, newUser.Email, result.Email)
		assert.Equal(t, newUser.Name, result.Name)
		assert.Equal(t, newUser.Password, result.Password)
		assert.Equal(t, newUser.Role, result.Role)
		assert.Equal(t, newUser.Status, result.Status)
	})

	t.Run("Email Has Been Registered", func(t *testing.T) {
		data.On("GetByEmail", newUser.Email).Return(&newUser, nil).Once()

		result, err := service.Register(newUser)

		assert.Error(t, err)
		assert.EqualError(t, err, "Email already registered by another user")
		assert.Nil(t, result)
	})

	t.Run("Hash Password Error", func(t *testing.T) {
		data.On("GetByEmail", newUser.Email).Return(&newUser, errors.New("")).Once()
		enkrip.On("HashPassword", newUser.Password).Return("", errors.New("Hash Password Error")).Once()

		result, err := service.Register(newUser)

		assert.Error(t, err)
		assert.EqualError(t, err, "Hash Password Error")
		assert.Nil(t, result)
	})

	t.Run("Failed Register", func(t *testing.T) {
		data.On("GetByEmail", newUser.Email).Return(&newUser, errors.New("")).Once()
		enkrip.On("HashPassword", newUser.Password).Return(newUser.Password, nil).Once()
		data.On("Register", newUser).Return(nil, errors.New("Failed to Register")).Once()

		result, err := service.Register(newUser)

		assert.Error(t, err)
		assert.EqualError(t, err, "Failed to Register")
		assert.Nil(t, result)
	})
}

func TestLogin(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)
	userData := users.User{
		Name:     "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "vanhau123",
	}

	t.Run("Success Login", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": uint(0), "role": "mockToken"}
		data.On("Login", userData.Email, userData.Password).Return(&userData, nil).Once()
		jwt.On("GenerateJWT", uint(0), "", "").Return(jwtResult).Once()
		result, err := service.Login(userData.Email, userData.Password)

		data.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userData.Email, result.Email)
		assert.Equal(t, "Irvan Hauwerich", result.Name)
		assert.Equal(t, jwtResult, result.Access)
	})

	t.Run("Password Incorrect", func(t *testing.T) {
		userFail := users.User{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("Login", userFail.Email, userFail.Password).Return(nil, errors.New("Incorrect Password")).Once()

		result, err := service.Login(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Incorrect Password")
		assert.Nil(t, result)
	})

	t.Run("Not Found", func(t *testing.T) {
		userFail := users.User{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("Login", userFail.Email, userFail.Password).Return(nil, errors.New("User Not Found / User Inactive")).Once()

		result, err := service.Login(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "User Not Found / User Inactive")
		assert.Nil(t, result)
	})

	t.Run("Server Error", func(t *testing.T) {
		userFail := users.User{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("Login", userFail.Email, userFail.Password).Return(nil, errors.New("Process Failed")).Once()

		result, err := service.Login(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Process Failed")
		assert.Nil(t, result)
	})

	t.Run("Token Failed", func(t *testing.T) {
		data.On("Login", userData.Email, userData.Password).Return(&userData, nil).Once()
		jwt.On("GenerateJWT", uint(0), "", "").Return(nil).Once()
		result, err := service.Login(userData.Email, userData.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Token Process Failed")
		assert.Nil(t, result)
	})
}

func TestGenerateJWT(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)
	userData := users.User{
		Name:     "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "vanhau123",
	}

	t.Run("Success Generate", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": uint(0), "role": "mockToken"}
		jwt.On("GenerateJWT", uint(0), "", "").Return(jwtResult).Once()
		data.On("GetByEmail", "irvanhau@gmail.com").Return(&userData, nil).Once()

		result, err := service.GenerateJwt("irvanhau@gmail.com")

		data.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByEmail", "irvanhau@gmail.com").Return(nil, errors.New("Process Failed")).Once()

		result, err := service.GenerateJwt("irvanhau@gmail.com")

		assert.Error(t, err)
		assert.EqualError(t, err, "Process Failed")
		assert.Nil(t, result)
	})

	t.Run("Token Failed", func(t *testing.T) {
		data.On("GetByEmail", "irvanhau@gmail.com").Return(&userData, nil).Once()
		jwt.On("GenerateJWT", uint(0), "", "").Return(nil).Once()

		result, err := service.GenerateJwt("irvanhau@gmail.com")

		assert.Error(t, err)
		assert.EqualError(t, err, "Token Process Failed")
		assert.Nil(t, result)
	})
}

func TestTokenResetVerify(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)
	userReset := users.UserResetPass{
		Email:     "irvanhau@gmail.com",
		Code:      "codetesting",
		ExpiresAt: time.Now().Add(time.Hour * 1),
	}

	t.Run("Success Reset", func(t *testing.T) {
		data.On("GetByCode", userReset.Code).Return(&userReset, nil).Once()

		result, err := service.TokenResetVerify(userReset.Code)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Input Error", func(t *testing.T) {
		data.On("GetByCode", userReset.Code).Return(nil, errors.New("Failed to verify token")).Once()

		result, err := service.TokenResetVerify(userReset.Code)

		assert.Error(t, err)
		assert.EqualError(t, err, "Failed to verify token")
		assert.Nil(t, result)
	})

	t.Run("Expire Time", func(t *testing.T) {
		userResetFail := users.UserResetPass{
			Email:     "irvanhau@gmail.com",
			Code:      "codetesting",
			ExpiresAt: time.Now().Add(time.Hour * -1),
		}
		data.On("GetByCode", userResetFail.Code).Return(&userResetFail, nil).Once()

		result, err := service.TokenResetVerify(userResetFail.Code)

		assert.Error(t, err)
		assert.EqualError(t, err, "token expired")
		assert.Nil(t, result)
	})
}

func TestForgetPasswordWeb(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)
	userData := users.User{
		Name:     "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "vanhau123",
		Role:     "Doctor",
	}

	t.Run("Success Forget Password", func(t *testing.T) {
		data.On("GetByEmail", userData.Email).Return(&userData, nil).Once()
		data.On("InsertCode", userData.Email, "code").Return(nil).Once()
		email.On("HTMLBody", userData.Role, userData.Name).Return("header", "htmlBody", "code").Once()
		email.On("SendEmail", userData.Email, "header", "htmlBody").Return(nil).Once()

		err := service.ForgetPasswordWeb(userData.Email)

		assert.Nil(t, err)
		data.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("Insert Code Error", func(t *testing.T) {
		data.On("GetByEmail", userData.Email).Return(&userData, nil).Once()
		email.On("HTMLBody", userData.Role, userData.Name).Return("header", "htmlBody", "code").Once()
		data.On("InsertCode", userData.Email, "code").Return(errors.New("Insert Code Failed")).Once()

		err := service.ForgetPasswordWeb(userData.Email)

		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Code Failed")
	})

	t.Run("Email Not Found", func(t *testing.T) {
		data.On("GetByEmail", userData.Email).Return(nil, errors.New("data not found")).Once()

		err := service.ForgetPasswordWeb(userData.Email)

		assert.Error(t, err)
		assert.EqualError(t, err, "data not found")
	})

	t.Run("Process Failed", func(t *testing.T) {
		data.On("GetByEmail", userData.Email).Return(nil, errors.New("process failed")).Once()

		err := service.ForgetPasswordWeb(userData.Email)

		assert.Error(t, err)
		assert.EqualError(t, err, "process failed")
	})

	t.Run("Send Email Error", func(t *testing.T) {
		data.On("GetByEmail", userData.Email).Return(&userData, nil).Once()
		data.On("InsertCode", userData.Email, "code").Return(nil).Once()
		email.On("HTMLBody", userData.Role, userData.Name).Return("header", "htmlBody", "code").Once()
		email.On("SendEmail", userData.Email, "header", "htmlBody").Return(errors.New("Send Email Error")).Once()

		err := service.ForgetPasswordWeb(userData.Email)

		assert.Error(t, err)
		assert.EqualError(t, err, "Send Email Error")
	})

}

func TestResetPassword(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)

	t.Run("Success Reset Password", func(t *testing.T) {
		enkrip.On("HashPassword", "password").Return("password", nil).Once()
		data.On("ResetPassword", "code", "email", "password").Return(nil).Once()

		err := service.ResetPassword("code", "email", "password")

		assert.Nil(t, err)
		data.AssertExpectations(t)
		enkrip.AssertExpectations(t)
	})

	t.Run("Hash Password Error", func(t *testing.T) {
		enkrip.On("HashPassword", "password").Return("", errors.New("Hash Password Error")).Once()

		err := service.ResetPassword("code", "email", "password")

		assert.Error(t, err)
		assert.EqualError(t, err, "Hash Password Error")
	})

	t.Run("Reset Password Error", func(t *testing.T) {
		enkrip.On("HashPassword", "password").Return("password", nil).Once()
		data.On("ResetPassword", "code", "email", "password").Return(errors.New("Reset Password Process Failed")).Once()

		err := service.ResetPassword("code", "email", "password")

		assert.Error(t, err)
		assert.EqualError(t, err, "Reset Password Process Failed")
	})
}

func TestUpdateProfile(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	enkrip := helper.NewHashInterface(t)
	data := mocks.NewUserDataInterface(t)
	email := helper.NewEmailInterface(t)
	service := New(data, jwt, email, enkrip)
	userProfile := users.UpdateProfile{
		Name:     "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "password",
	}

	t.Run("Hash Password Error", func(t *testing.T) {
		enkrip.On("HashPassword", userProfile.Password).Return("", errors.New("Hash Password Error")).Once()

		res, err := service.UpdateProfile(1, userProfile)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Hash Password Error")
	})

	t.Run("Update Profile Failed", func(t *testing.T) {
		enkrip.On("HashPassword", userProfile.Password).Return(userProfile.Password, nil).Once()
		data.On("UpdateProfile", 1, userProfile).Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdateProfile(1, userProfile)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update Profile", func(t *testing.T) {
		enkrip.On("HashPassword", userProfile.Password).Return(userProfile.Password, nil).Once()
		data.On("UpdateProfile", 1, userProfile).Return(true, nil).Once()

		res, err := service.UpdateProfile(1, userProfile)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		enkrip.AssertExpectations(t)
		data.AssertExpectations(t)
	})
}
