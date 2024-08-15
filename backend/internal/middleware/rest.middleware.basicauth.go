package middleware

import (
	"encoding/base64"
	"strings"

	"fudjie.waizly/backend-test/internal/helper"

	"github.com/gofiber/fiber/v2"
)

func (m *RestMiddlewareModule) GuardBasicAuthentication(fc *fiber.Ctx) error {
	var (
		err error

		authKey = fc.Get("Authorization")
	)

	if !m.authConfig.EnableBasicAuth {
		return fc.Next()
	}

	if authKey == "" {
		return helper.NewUnauthorizedErr(err, "incorrect username and password")
	}

	_, err = m.validateAuthorization(authKey)

	if err != nil {
		return helper.NewUnauthorizedErr(err, "incorrect username and password")
	}

	return fc.Next()
}

func (m *RestMiddlewareModule) validateAuthorization(key string) (bool, error) {
	var (
		err        error
		encodedb64 string
		decodedb64 []byte
		res        []string
	)

	encodedb64 = strings.Split(key, "Basic ")[1]

	decodedb64, err = base64.StdEncoding.DecodeString(encodedb64)

	if err != nil {
		return false, err
	}

	res = strings.Split(string(decodedb64), ":")

	if res[0] == m.authConfig.Username && res[1] == m.authConfig.Password {
		return true, nil
	}

	return false, nil
}
