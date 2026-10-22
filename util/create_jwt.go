package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UlrEncode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payLoadB64 := base64UlrEncode(byteArrData)

	message := headerB64 + "." + payLoadB64

	byteArrSecret := []byte(secret)
	byteArrMEssage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMEssage)

	signature := h.Sum(nil)
	signatureB64 := base64UlrEncode(signature)

	jwt := headerB64 + "." + payLoadB64 + "." + signatureB64

	return jwt, nil
}

func base64UlrEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
