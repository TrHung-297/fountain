/* !!
 * File: crypto_util.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:30:19 am
 
 */

package crypto

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

func Sha256Digest(data []byte) []byte {
	r := sha256.Sum256(data)
	return r[:]
}

func Sha1Digest(data []byte) []byte {
	r := sha1.Sum(data)
	return r[:]
}

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateStringNonce(size int) string {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func CalculatePaddingSize(lenData int32) int32 {
	var additionalSize = (32 + lenData) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}
	if additionalSize < 12 {
		additionalSize += 16
	}
	return additionalSize
}

func CalculateSha1ForNotify(lenData, additionalSize int32, secret []byte) []byte {
	bytesCheck := make([]byte, 24+lenData+additionalSize)
	h := sha1.New()
	h.Write(secret)
	bs := h.Sum(nil)
	copy(bytesCheck[:8], bs[len(bs)-8:])
	return bytesCheck
}

func CalculateMsgKey(customNotifyJSON, bytesCheck []byte, secret []byte) []byte {
	messageKey := make([]byte, 32)
	t_d := make([]byte, 0, 32+len(customNotifyJSON))
	t_d = append(t_d, secret[88+8:88+8+32]...)
	t_d = append(t_d, customNotifyJSON...)
	copy(messageKey, Sha256Digest(t_d))
	msgKey := messageKey[8 : 8+16]
	copy(bytesCheck[8:8+16], msgKey)
	return msgKey
}

func CalculateAESKeyAndAESIV(msgKey, secret []byte) ([]byte, []byte) {
	var x = 8
	t_a := make([]byte, 0, 52)
	t_a = append(t_a, msgKey[:16]...)
	t_a = append(t_a, secret[x:x+36]...)
	sha256_a := Sha256Digest(t_a)

	t_b := make([]byte, 0, 52)
	t_b = append(t_b, secret[40+x:40+x+36]...)
	t_b = append(t_b, msgKey[:16]...)
	sha256_b := Sha256Digest(t_b)

	aesKey := make([]byte, 0, 32)
	aesKey = append(aesKey, sha256_a[:8]...)
	aesKey = append(aesKey, sha256_b[8:8+16]...)
	aesKey = append(aesKey, sha256_a[24:24+8]...)

	aesIV := make([]byte, 0, 32)
	aesIV = append(aesIV, sha256_b[:8]...)
	aesIV = append(aesIV, sha256_a[8:8+16]...)
	aesIV = append(aesIV, sha256_b[24:24+8]...)
	return aesKey, aesIV
}

func AddLenData(customNotifyJSON []byte) []byte {
	lenCustomNotifyJSONByteArr := make([]byte, 4)
	binary.LittleEndian.PutUint32(lenCustomNotifyJSONByteArr, uint32(len(customNotifyJSON)))
	customNotifyJSON = append(lenCustomNotifyJSONByteArr, customNotifyJSON...)
	return customNotifyJSON
}
