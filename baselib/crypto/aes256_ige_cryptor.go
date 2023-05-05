/* !!
 * File: aes256_ige_cryptor.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:30:14 am
 
 */

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

var (
	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("AES256IGE: invalid blocksize - data too small to encrypt")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("AES256IGE: invalid padding - data not divisible by block size")
)

// AES256IGECryptor type
type AES256IGECryptor struct {
	aesKey []byte
	aesIV  []byte
}

// NewAES256IGECryptor func
func NewAES256IGECryptor(aesKey, aesIV []byte) *AES256IGECryptor {
	// guard conditions
	if (len(aesIV)) < aes.BlockSize {
		return nil
	}

	k := len(aesKey)
	switch k {
	default:
		return nil
	case 16, 24, 32:
		break
	}
	return &AES256IGECryptor{
		aesKey: aesKey,
		aesIV:  aesIV,
	}
}

// Encrypt func
// The data length must be a multiple of aes.BlockSize(16). If not, please call the caller.
func (c *AES256IGECryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, c.aesIV[:aes.BlockSize])
	copy(y, c.aesIV[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

// Decrypt func
func (c *AES256IGECryptor) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, c.aesIV[:aes.BlockSize])
	copy(y, c.aesIV[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil
}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}

// -------------------------- SIMPLE --------------------------

func pkcs7AddPadding(data []byte, blocksize int) []byte {
	n := blocksize - (len(data) % blocksize)
	pb := make([]byte, len(data)+n)
	copy(pb, data)
	copy(pb[len(data):], bytes.Repeat([]byte{byte(n)}, n))
	return pb
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7UnPadding(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}

	if len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}

	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}

	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}

	return b[:len(b)-n], nil
}

// EncryptSimpleWithBase64 func
// The data length must be a multiple of aes.BlockSize(16). If not, please call the caller.
// The result was encoded by base64
func (c *AES256IGECryptor) EncryptSimpleWithBase64(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}

	i := len(data)
	if i/block.BlockSize() != 0 {
		data = append(data, make([]byte, block.BlockSize()-i%block.BlockSize())...)
	}

	content := []byte(data)
	content = pkcs7AddPadding(content, block.BlockSize())
	cipher.NewCBCEncrypter(block, c.aesIV).CryptBlocks(content, content)

	result := base64.StdEncoding.EncodeToString([]byte(content))
	return []byte(result), nil
}

// DecryptSimpleWithBase64 func
func (c *AES256IGECryptor) DecryptSimpleWithBase64(base64Data string) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(base64Data)

	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, ErrInvalidBlockSize
	}

	if len(data)%aes.BlockSize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	cipher.NewCBCDecrypter(block, c.aesIV).CryptBlocks(data, data)
	text, err := pkcs7UnPadding(data, aes.BlockSize)
	if err != nil {
		return []byte{}, err
	}
	return text, nil
}
