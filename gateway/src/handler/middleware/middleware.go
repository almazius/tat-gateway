package middleware

import (
	"gateway/gateway/config"
	interfaces "gateway/gateway/pkg/interface"
	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	authService interfaces.AuthService
}

func NewMiddleware(auth interfaces.AuthService) *middleware {
	return &middleware{authService: auth}
}

func (m *middleware) Authed() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !config.C().AuthServer.Enabled {
			return c.Next()
		}

		// some logic...

		return c.Next()
	}
}
