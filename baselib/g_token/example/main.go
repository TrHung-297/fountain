/* !!
 * File: main.go
 * File Created: Friday, 19th November 2021 6:18:27 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Friday, 19th November 2021 6:18:36 pm
 
 */

package main

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	bearToken := `eyJhbGciOiJSUzI1NiIsImtpZCI6ImY1NWUyOTRlZWRjMTY3Y2Q5N2JiNWE4MTliYmY3OTA2MzZmMTIzN2UiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiUGhhbSBWYW4gS2ltIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hLS9BT2gxNEdncW0zaDhoZTBCOGtrYWRDSnBrczdKdF9tLUpkb1JJdFhWZ1hjPXM5Ni1jIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL3BsYXktYXV0aC01NmEyMyIsImF1ZCI6InBsYXktYXV0aC01NmEyMyIsImF1dGhfdGltZSI6MTYzNzMxNzMyNywidXNlcl9pZCI6Ikdod29DSDNGN1VPSGVPY2lURHJNNWhNMWk1cjIiLCJzdWIiOiJHaHdvQ0gzRjdVT0hlT2NpVERyTTVoTTFpNXIyIiwiaWF0IjoxNjM3MzE3MzI3LCJleHAiOjE2MzczMjA5MjcsImVtYWlsIjoia2ltcHZAZ3R2LnZuIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMTQ3Nzg4Nzg5NTkyMDE5NDQ2NTgiXSwiZW1haWwiOlsia2ltcHZAZ3R2LnZuIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.q47GrCAdMUWB_ApzmoOTNQeu5WXX1Ndr6AFzSyS51BgO0ZiuTaMdLGjbh4l860CRptguSMs8mh4BSaKV16PNxFtWF_TZCOU2L6DJaWYpnb6UOxFlOeXhdeZxQ30sjSvvmERgKgUPtx_GZ4DU2TiJQOA5LP72ZXkrIghENj4Zja8vOPvEUq29G3m1hWkTwnjOc1P6imqcMcWJ7efoEe79_iaTYnXBaeEPbYk0qdMO7aBJzzfMFJ63kSgYaiBKijATxiRBrKHDH0Mp4ouEUax_QCe3tWOOZH40JE42mlpjgZOap5pS_yzPj9KRCLFGAoHSZF-92DV2mS1dc7YJKgy56g`
	token, err := jwt.Parse(bearToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Signing Key
		signKey := "okedfasdfasdfasdfala"
		return []byte(signKey), nil
	})

	if err != nil {
		panic(err)
	}

	// Check Token Valid Expire
	if token == nil {
		err = errors.New("not found or can not parse token")
		return
	}

	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessToken, ok := claims["KTokenAccessUUIDKey"].(string)
		if !ok {
			err = errors.New("token Invalid or Expire - not found access uuid key")
			return
		}
		_ = accessToken
	}
}
