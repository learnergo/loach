package model

type Client interface {
	GetAdmin() (*Identity, error)

	Register(*RegisterRequest) (*RegisterResponse, error)
	Enroll(*EnrollRequest) (*EnrollResponse, error)
}
