package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type Credentials struct {
	Key        string `json:"api_key"`
	SecretKey  string `json:"api_secret_key"`
	Passphrase string
	Type       string // read, read-trade, read-withdraw
	UserID     string `json:"user_id"`
}

/*
 signing a message
 using: hmac sha256 + base64
  eg:
    message = Pre_hash function comment
    secretKey = E65791902180E9EF4510DB6A77F6EBAE
  return signed string = TO6uwdqz+31SIPkd4I+9NiZGmVH74dXi+Fd5X0EzzSQ=
*/
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}
