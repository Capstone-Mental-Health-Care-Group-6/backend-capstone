// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"
)

// JWTInterface is an autogenerated mock type for the JWTInterface type
type JWTInterface struct {
	mock.Mock
}

// CheckID provides a mock function with given fields: c
func (_m *JWTInterface) CheckID(c echo.Context) interface{} {
	ret := _m.Called(c)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(echo.Context) interface{}); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// CheckRole provides a mock function with given fields: c
func (_m *JWTInterface) CheckRole(c echo.Context) interface{} {
	ret := _m.Called(c)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(echo.Context) interface{}); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// ExtractToken provides a mock function with given fields: token
func (_m *JWTInterface) ExtractToken(token *jwt.Token) map[string]interface{} {
	ret := _m.Called(token)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(*jwt.Token) map[string]interface{}); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GenerateJWT provides a mock function with given fields: userID, role, status
func (_m *JWTInterface) GenerateJWT(userID uint, role string, status string) map[string]interface{} {
	ret := _m.Called(userID, role, status)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(uint, string, string) map[string]interface{}); ok {
		r0 = rf(userID, role, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id, role, status
func (_m *JWTInterface) GenerateToken(id uint, role string, status string) string {
	ret := _m.Called(id, role, status)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, string, string) string); ok {
		r0 = rf(id, role, status)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetID provides a mock function with given fields: c
func (_m *JWTInterface) GetID(c echo.Context) (uint, error) {
	ret := _m.Called(c)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context) (uint, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(echo.Context) uint); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(echo.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshJWT provides a mock function with given fields: accessToken, refreshToken
func (_m *JWTInterface) RefreshJWT(accessToken string, refreshToken *jwt.Token) (map[string]interface{}, error) {
	ret := _m.Called(accessToken, refreshToken)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *jwt.Token) (map[string]interface{}, error)); ok {
		return rf(accessToken, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(string, *jwt.Token) map[string]interface{}); ok {
		r0 = rf(accessToken, refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, *jwt.Token) error); ok {
		r1 = rf(accessToken, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: token
func (_m *JWTInterface) ValidateToken(token string) (*jwt.Token, error) {
	ret := _m.Called(token)

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTInterface creates a new instance of JWTInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTInterface {
	mock := &JWTInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
