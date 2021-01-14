package config

import "github.com/gbrlsnchs/jwt/v3"

// AuthConfiguration describes Auth options in config.yml
type AuthConfiguration struct {
	Header     string
	Prefix     string
	PrivateKey string
	Algorithm  jwt.Algorithm
}
