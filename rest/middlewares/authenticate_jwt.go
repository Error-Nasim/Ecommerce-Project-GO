package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerece/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		cnf := config.GetConfig()
		byteArrSecret := []byte(cnf.JwtSecretKey)
		byteArrMEssage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMEssage)

		hash := h.Sum(nil)
		newSignature := base64UlrEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func base64UlrEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
