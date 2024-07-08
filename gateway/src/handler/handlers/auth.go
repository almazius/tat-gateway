package handlers

import (
	"fmt"
	interfaces "gateway/gateway/pkg/interface"
	"gateway/gateway/pkg/validator"
	"gateway/gateway/src/handler/models"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type LoginHandler struct {
	authService interfaces.AuthService
}

func NewLoginHandler(authService interfaces.AuthService) *LoginHandler {
	return &LoginHandler{authService: authService}
}

func (l *LoginHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginInfo := new(models.LoginRequestDTO)
		err := validator.ParseBody(c, loginInfo)
		if err != nil {
			return fmt.Errorf("failed to parse login info: %w", err)
		}

		loginInfo.UniqueInformation = map[string]string{
			"User-Agent": c.Get("User-Agent"), // get headers name from config
		}

		userInfo, err := l.authService.Login(loginInfo)
		if err != nil {
			slog.ErrorContext(c.UserContext(), "failed to login",
				slog.Any("error", err))
			return c.Status(fiber.StatusForbidden).JSON(models.ErrorMessage{
				Message: "login or password is not actual",
			})
		}

		c.Cookie(&fiber.Cookie{Name: "session_id", Value: userInfo.SessionId.String()}) // todo use consts

		return c.SendStatus(fiber.StatusOK)
	}
}

func (l *LoginHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		registerInfo := new(models.RegisterRequestDTO)
		err := validator.ParseBody(c, registerInfo)
		if err != nil {
			return fmt.Errorf("failed to parse register info: %w", err)
		}

		registerInfo.UniqueInformation = map[string]string{
			"User-Agent": c.Get("User-Agent"), // get headers name from config
		}

		userInfo, err := l.authService.Register(registerInfo)
		if err != nil {
			slog.ErrorContext(c.UserContext(), "failed to register",
				slog.Any("error", err))
			return c.Status(fiber.StatusForbidden).JSON(models.ErrorMessage{
				Message: "failed to register", // todo add normal description on error
			})
		}

		c.Cookie(&fiber.Cookie{Name: "session_id", Value: userInfo.SessionId.String()}) // todo use consts

		return c.SendStatus(fiber.StatusOK)
	}
}
