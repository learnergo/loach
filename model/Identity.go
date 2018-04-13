package model

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
)

type Identity struct {
	Key  interface{}
	Cert *x509.Certificate
}

func MarshalIdentity(identity *Identity) (string, error) {

	var key, cert string
	switch identity.Key.(type) {
	case *ecdsa.PrivateKey:
		cast := identity.Key.(*ecdsa.PrivateKey)
		b, err := x509.MarshalECPrivateKey(cast)
		if err != nil {
			return "", err
		}
		block := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
		key = base64.StdEncoding.EncodeToString(block)

	default:
		return "", errors.New("Error peivate key Type")
	}

	cert = CertToString(identity.Cert)
	str, err := json.Marshal(map[string]string{"cert": cert, "key": key})
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func CertToString(cert *x509.Certificate) string {
	return base64.StdEncoding.EncodeToString(cert.Raw)
}
