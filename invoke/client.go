package invoke

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/learnergo/loach/config"
	"github.com/learnergo/loach/crypto"

	"github.com/learnergo/loach/model"
)

type ClientImpl struct {
	Url        string
	ServerName string
	Crypto     crypto.Crypto
	Config     config.ClientConfig
}

func (client *ClientImpl) GetAdmin() (*model.Identity, error) {
	return loadIdentity(client.Config.AdminKey, client.Config.AdminCert)
}

func (client *ClientImpl) GetServer() (string, string) {
	return client.Url, client.ServerName
}

func (client *ClientImpl) Register(request *model.RegisterRequest) (*model.RegisterResponse, error) {
	return register(client, request)
}

func (client *ClientImpl) Enroll(request *model.EnrollRequest) (*model.EnrollResponse, error) {
	return enroll(client, request)
}

func (client *ClientImpl) createAuthToken(identity *model.Identity, request []byte) (string, error) {

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

func (client *ClientImpl) getTransport() *http.Transport {
	var tr *http.Transport
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return tr
}

func loadIdentity(keyPath string, certPath string) (*model.Identity, error) {
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	pCert, _ := pem.Decode(certData)
	pKey, _ := pem.Decode(keyData)

	cert, err := x509.ParseCertificate(pCert.Bytes)
	if err != nil {
		return nil, err
	}
	key, err := x509.ParsePKCS8PrivateKey(pKey.Bytes)
	if err != nil {
		return nil, err
	}
	return &model.Identity{Cert: cert, Key: key}, nil
}
