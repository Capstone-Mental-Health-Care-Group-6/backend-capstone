package helper

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type JWTInterface interface {
	GenerateJWT(userID uint, role, status string) map[string]any
	GenerateToken(id uint, role, status string) string
	ExtractToken(token *jwt.Token) map[string]interface{}
	ValidateToken(token string) (*jwt.Token, error)
	GetID(c echo.Context) (uint, error)
	CheckRole(c echo.Context) interface{}
	CheckID(c echo.Context) interface{}
	RefreshJWT(accessToken string, refreshToken *jwt.Token) (map[string]any, error)
}

type JWT struct {
	signKey    string
	refreshKey string
}

func New(signKey, refreshKey string) JWTInterface {
	return &JWT{
		signKey:    signKey,
		refreshKey: refreshKey,
	}
}

func (j *JWT) GenerateJWT(userID uint, role, status string) map[string]any {
	var result = map[string]any{}
	var accessToken = j.GenerateToken(userID, role, status)
	var refreshToken = j.generateRefreshToken(accessToken)
	if accessToken == "" {
		return nil
	}
	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	return result
}

func (j *JWT) RefreshJWT(accessToken string, refreshToken *jwt.Token) (map[string]any, error) {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()

	if err != nil {
		logrus.Error("Get Token Expiration Error, ", err.Error())
		return nil, errors.New("Token Expiration Error")
	}

	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newClaim = jwt.MapClaims{}
		newToken, err := jwt.ParseWithClaims(accessToken, newClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.signKey), nil
		})

		if err != nil {
			log.Error(err.Error())
			return nil, errors.New(err.Error())
		}

		newClaim = newToken.Claims.(jwt.MapClaims)
		newClaim["iat"] = time.Now().Unix()
		newClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

		result["access_token"] = newToken.Raw
		result["refresh_token"] = newSignedRefreshToken

		return result, nil
	}
	return nil, errors.New("Refresh Token Not Valid && Expired")
}

func (j *JWT) GenerateToken(id uint, role, status string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["status"] = status
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) generateRefreshToken(accessToken string) string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(j.refreshKey))

	if err != nil {
		return ""
	}

	return refreshToken
}
func (j *JWT) ExtractToken(token *jwt.Token) map[string]interface{} {
	if token.Valid {
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		if expTime.Time.Compare(time.Now()) > 0 {
			var mapClaim = claims.(jwt.MapClaims)
			var result = map[string]interface{}{}
			result["id"] = mapClaim["id"]
			result["role"] = mapClaim["role"]
			result["status"] = mapClaim["status"]
			return result
		}
		logrus.Error("Token Expired")
		return nil
	}
	return nil
}

func (j *JWT) ValidateToken(token string) (*jwt.Token, error) {
	var authHeader = token[7:]
	parsedToken, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.signKey), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func (j *JWT) GetID(c echo.Context) (uint, error) {
	authHeader := c.Request().Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return 0, err
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	idFloat, ok := mapClaim["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("ID not found or not a valid number")
	}

	idUint := uint(idFloat)
	return idUint, nil
}
func (j *JWT) CheckRole(c echo.Context) interface{} {
	authHeader := c.Request().Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return c.JSON(http.StatusUnauthorized, FormatResponse("Token is not valid", nil))
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	role := mapClaim["role"]

	return role
}

func (j *JWT) CheckID(c echo.Context) any {
	authHeader := c.Request().Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return c.JSON(http.StatusUnauthorized, FormatResponse("Token is not valid", nil))
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	id := mapClaim["id"]

	return id
}
