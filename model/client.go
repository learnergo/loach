package model

type Client interface {
	Register(*Identity, *RegisterRequest) (*RegisterResponse, error)
	Enroll(*EnrollRequest) (*EnrollResponse, error)
}
