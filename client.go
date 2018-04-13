package loach

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/learnergo/loach/config"
	"github.com/learnergo/loach/crypto"
	"github.com/learnergo/loach/invoke"
	"github.com/learnergo/loach/model"
)

func NewClient(path string) (model.Client, error) {

	config, err := config.NewClientConfig(path)
	if err != nil {
		return nil, err
	}

	c, err := getCrypto(config.CryptoConfig)
	if err != nil {
		return nil, err
	}

	return &invoke.ClientImpl{
		Url:    config.Url,
		Crypto: c,
		Config: *config,
	}, nil
}

func LoadIdentity(keyPath string, certPath string) (*model.Identity, error) {
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

func getCrypto(cc config.CryptoConfig) (crypto.Crypto, error) {
	switch cc.Family {
	case "ecdsa":
		c, err := crypto.NewCrypto(cc)
		if err != nil {
			return nil, err
		}
		return c, nil
	default:
		return nil, errors.New("Error Crypto")
	}
}
