package interfaces

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	Authed() fiber.Handler
	OTPCheck() fiber.Handler

	SellerCheck() fiber.Handler
}
