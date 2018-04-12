package loach

import (
	"errors"

	"github.com/learnergo/loach/config"
	"github.com/learnergo/loach/crypto"
	"github.com/learnergo/loach/invoke"
)

func NewClient(path string) (*invoke.Client, error) {

	config, err := config.NewClientConfig(path)
	if err != nil {
		return nil, err
	}

	var c crypto.Crypto
	switch config.CryptoConfig.Family {
	case "ecdsa":
		c, err = crypto.NewCrypto(config.CryptoConfig)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Error Crypto")
	}

	return &invoke.Client{
		Url:    config.Url,
		Crypto: c,
		Config: *config,
	}, nil
}
