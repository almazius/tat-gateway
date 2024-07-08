package interfaces

import "gateway/gateway/src/handler/models"

type AuthService interface {
	Login(loginInfo *models.LoginRequestDTO) (models.LoginResponseDTO, error)
	Register(registerInfo *models.RegisterRequestDTO) (models.LoginResponseDTO, error)

	// some methods...
}
