package handler

import (
	interfaces "gateway/gateway/pkg/interface"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) initRoutes() {
	api := s.app.Group("/api")

	auth := api.Group("/auth")
	s.loginHandlers(auth)

	market := api.Group("/market")
	s.marketService(market, s.mw)

	user := api.Group("/user_service")
	s.userService(user, s.mw)

	billing := api.Group("/billing")
	s.billingService(billing, s.mw)
}

func (s *Server) loginHandlers(router fiber.Router) {
	router.Post("/login", s.LoginHandlers.Login())
	router.Post("/register", s.LoginHandlers.Register())

	router.Post("/reset_password", nil)
	router.Post("/verify_code", nil)
}

func (s *Server) userService(router fiber.Router, mw interfaces.Middleware) {
	router.Get("/self_info", mw.Authed(), nil)

	router.Patch("/update_password", mw.Authed(), mw.OTPCheck(), nil)
	router.Patch("/update_otp", mw.Authed(), mw.OTPCheck(), nil)
}

func (s *Server) marketService(router fiber.Router, mw interfaces.Middleware) {
	// get all offers
	router.Get("/", nil)
	router.Get("/details/:id/free_records", nil)
	router.Get("/details/:id/feedbacks", nil)
	router.Get("/details/:id", nil)

	// sellers info and offers
	router.Get("/seller/:id/offers")
	router.Get("/seller/:id/feedbacks")
	router.Get("/seller/:id")

	router.Post("/subscribe", mw.Authed(), nil)
	router.Post("/unsubscribe", mw.Authed(), nil)

	router.Post("/feedback", mw.Authed(), nil)

	// пожаловаться на объявление
	router.Post("/report", mw.Authed(), nil)

	seller := router.Group("/seller_panel", mw.SellerCheck())
	seller.Post("/offer", nil)
	seller.Patch("/offer/:offer_id", nil)
	seller.Delete("/offer/:offer_id", nil)
}

func (s *Server) billingService(router fiber.Router, mw interfaces.Middleware) {
	router.Post("/payment", mw.Authed(), nil)
	router.Get("/payment_history", mw.Authed(), nil)
}
