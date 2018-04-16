package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	Url          string `yaml:"url"`
	AdminKey     string `yaml:"admin_key"`
	AdminCert    string `yaml:"admin_cert"`
	CryptoConfig `yaml:"crypto"`
}

type CryptoConfig struct {
	Family    string `yaml:"family"`
	Algorithm string `yaml:"algorithm"`
	Hash      string `yaml:"hash"`
}

func NewClientConfig(path string) (*ClientConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := new(ClientConfig)
	err = yaml.Unmarshal([]byte(data), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
