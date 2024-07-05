package handler

import (
	interfaces "gateway/gateway/pkg/interface"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) initRoutes() {
	api := s.app.Group("/api")

	auth := api.Group("/auth")
	loginHandlers(auth)

}

func loginHandlers(router fiber.Router) {
	router.Post("/login", nil)
	router.Post("/register", nil)

	router.Post("/reset_password", nil)
	router.Post("/verify_code", nil)
}

func userService(router fiber.Router, mw interfaces.Middleware) {
	router.Get("/user_info", mw.Authed(), nil)

	router.Patch("/update_password", mw.Authed(), mw.OTPCheck(), nil)
	router.Patch("/update_otp", mw.Authed(), mw.OTPCheck(), nil)
}

func marketService(router fiber.Router, mw interfaces.Middleware) {
	// get all offers
	router.Get("/", nil)
	router.Get("/details/:id/free_records", nil)
	router.Get("/details/:id", nil)

	router.Post("/subscribe", mw.Authed(), nil)
	router.Post("/unsubscribe", mw.Authed(), nil)

	seller := router.Group("/seller", mw.SellerCheck())
	seller.Post("/offer", nil)
	seller.Patch("/offer/:offer_id", nil)
	seller.Delete("/offer/:offer_id", nil)
}

func billingService(router fiber.Router, mw interfaces.Middleware) {
	router.Post("/payment", mw.Authed(), nil)
	router.Get("/payment_history", mw.Authed(), nil)
}
