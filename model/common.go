package model

import (
	"crypto/x509"
	"crypto/x509/pkix"
)

type Response struct {
	Success  bool          `json:"success"`
	Errors   []ResponseErr `json:"errors"`
	Messages []string      `json:"messages"`
}

type ResponseErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	EnrolmentId    string              `json:"id"`
	Type           string              `json:"type"`
	Secret         string              `json:"secret,omitempty"`
	MaxEnrollments int                 `json:"max_enrollments,omitempty"`
	Affiliation    string              `json:"affiliation"`
	Attrs          []RegisterAttribute `json:"attrs"`
	CAName         string              `json:"caname,omitempty"`
}

type RegisterAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	ECert bool   `json:"ecert,omitempty"`
}

type RegisterResponse struct {
	Response
	Result RegisterCredentialResponse `json:"result"`
}

type RegisterCredentialResponse struct {
	Secret string `json:"secret"`
}

type EnrollRequest struct {
	EnrollID string
	Secret   string
	Name     pkix.Name
	Profile  string            `json:"profile,omitempty"`
	Label    string            `json:"label,omitempty"`
	CAName   string            `json:"caname,omitempty"`
	Hosts    []string          `json:"hosts"`
	Attrs    []EnrollAttribute `json:"attr_reqs,omitempty"`
}

type EnrollAttribute struct {
	Name     string `json:"name"`
	Optional bool   `json:"optional,omitempty"`
}

type EnrollResponse struct {
	Identity  *Identity
	CertChain *x509.Certificate
}

type CreateCsrRequest struct {
	Name      pkix.Name
	Key       interface{}
	Hosts     []string
	Algorithm string
}
