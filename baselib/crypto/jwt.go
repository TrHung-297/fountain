/* !!
 * File: jwt.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:30:36 am
 
 */

package crypto

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	OvcTokenExpirationTime = 1
	OvcAppID               = 13213123
)

// OVCGenerateToken func
func OVCGenerateToken(roomID int64, userID int32) (string, error) {
	claims := jwt.MapClaims{}
	claims["app"] = OvcAppID
	claims["rid"] = strconv.FormatInt(roomID, 10)
	claims["uid"] = strconv.FormatInt(int64(userID), 10)
	claims["exp"] = time.Now().Add(time.Hour * OvcTokenExpirationTime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	at, err := token.SignedString([]byte("1q2dGu5pzikcrECJgW3ADfXX3EsmoD99SYvSVCpDsJrAqxou5tUNbHPvkEFI4bTS"))
	if err != nil {
		return "", err
	}

	return at, nil
}
