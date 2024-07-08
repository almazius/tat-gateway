package handler

import (
	"fmt"
	interfaces "gateway/gateway/pkg/interface"
	"gateway/gateway/src/handler/handlers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
	mw  interfaces.Middleware

	LoginHandlers handlers.LoginHandler
}

func NewServer() *Server {
	return &Server{
		app: fiber.New(fiber.Config{}),
	}
}

func (s *Server) ListenAndServe(host string, port int) error {
	err := s.app.Listen(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
