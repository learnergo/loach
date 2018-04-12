package invoke

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"

	"github.com/learnergo/loach/config"
	"github.com/learnergo/loach/crypto"

	"github.com/learnergo/loach/model"
)

type CaClient interface {
	Register(*model.Identity, *model.RegisterRequest) (string, error)
	Enroll(*model.EnrollRequest) (*model.EnrollResponse, error)
}

type Client struct {
	Url    string
	Crypto crypto.Crypto
	Config config.ClientConfig
}

func (client *Client) Register(admin *model.Identity, request *model.RegisterRequest) (*model.RegisterResponse, error) {
	return Register(client, admin, request)
}

func (client *Client) Enroll(request *model.EnrollRequest) (*model.EnrollResponse, error) {
	return Enroll(client, request)
}

func (client *Client) CreateAuthToken(identity *model.Identity, request []byte) (string, error) {

	encPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: identity.Cert.Raw})
	encCert := base64.StdEncoding.EncodeToString(encPem)
	body := base64.StdEncoding.EncodeToString(request)
	sigString := body + "." + encCert
	sig, err := client.Crypto.Sign([]byte(sigString), identity.Key)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", encCert, base64.StdEncoding.EncodeToString(sig)), nil
}

func (client *Client) GetTransport() *http.Transport {
	var tr *http.Transport
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return tr
}
