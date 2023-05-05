/* !!
 * File: aes_gtv_cryptor.go
 * File Created: Thursday, 15th July 2021 5:40:57 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 15th July 2021 5:40:57 pm
 
 */

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/TrHung-297/fountain/baselib/g_log"
)

type G_AES256Cryptor struct {
	SecretKey  []byte
	InitVector []byte
	BlockMode  cipher.BlockMode
}

const (
	KG_AESEncoderEnvKey string = "G_AESEncoder"
)

var gAES256Instance *G_AES256Cryptor

func GetG_AES256CryptorInstance() *G_AES256Cryptor {
	if gAES256Instance == nil {
		err := fmt.Errorf("need install GetG_AES256Cryptor first")
		g_log.WithError(err).Error(err)
		panic(err)
	}

	return gAES256Instance
}

func InstallG_AES256Cryptor(envKeys ...string) *G_AES256Cryptor {
	envKey := KG_AESEncoderEnvKey
	for _, key := range envKeys {
		if key != "" {
			envKey = key
			break
		}
	}

	getG_AESConfigFromEnv(envKey)
	if gAESConfig == nil || len(gAESConfig.SecretKey) == 0 {
		err := fmt.Errorf("not found config for G_AESEncoder from: %q", envKey)
		g_log.V(1).WithError(err).Errorf("getG_AESConfigFromEnv - Error: %+v", err)
		panic(err)
	}

	return newG_AES256Cryptor([]byte(gAESConfig.SecretKey), []byte(gAESConfig.InitVector))
}

func newG_AES256Cryptor(secretKey, initVector []byte) *G_AES256Cryptor {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		err := fmt.Errorf("newG_AES256Cryptor error: %+v", err)
		g_log.V(1).WithError(err).Errorf("newG_AES256Cryptor: %+v", err)
		panic(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, initVector)

	gAES256Instance = &G_AES256Cryptor{
		SecretKey:  secretKey,
		InitVector: initVector,
		BlockMode:  blockMode,
	}

	return gAES256Instance
}

// Encrypt

func (c *G_AES256Cryptor) Encrypt(message []byte) (encryptedMessage []byte, err error) {
	block, err := aes.NewCipher(c.SecretKey)
	if err != nil {
		return []byte(""), err
	}

	encryptedMessage = make([]byte, len(message))
	mode := cipher.NewCBCEncrypter(block, c.InitVector)
	mode.CryptBlocks(encryptedMessage, message)
	return encryptedMessage, nil
}

// EncryptSimpleWithBase64 func
// The data length must be a multiple of aes.BlockSize(16). If not, please call the caller.
// The result was encoded by base64
func (c *G_AES256Cryptor) EncryptSimpleWithBase64(message []byte) (string, error) {
	encryptedMessage, err := c.Encrypt(message)
	if err != nil {
		return "", err
	}

	base64Encrypted := base64.StdEncoding.EncodeToString(encryptedMessage)
	return base64Encrypted, nil
}

func (c *G_AES256Cryptor) DesEncryption(text []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.SecretKey)
	if err != nil {
		return nil, nil
	}

	blockMode := cipher.NewCBCEncrypter(block, c.InitVector)
	cipherText := make([]byte, len(text))
	blockMode.CryptBlocks(cipherText, text)
	cipherText = PKCS5Padding(cipherText, blockMode.BlockSize())
	return cipherText, nil
}

// EncryptSimpleWithBase64 func
// The data length must be a multiple of aes.BlockSize(16). If not, please call the caller.
// The result was encoded by base64
func (c *G_AES256Cryptor) DesEncryptSimpleWithBase64(message []byte) (string, error) {
	encryptedMessage, err := c.DesEncryption(message)
	if err != nil {
		return "", err
	}

	base64Encrypted := base64.StdEncoding.EncodeToString(encryptedMessage)
	return base64Encrypted, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// DeCrypt

func (c *G_AES256Cryptor) Decrypt(encryptedMessage []byte) (decryptedMessage []byte, err error) {
	block, err := aes.NewCipher(c.SecretKey)
	if err != nil {
		return []byte(""), err
	}

	decryptedMessage = make([]byte, len(encryptedMessage))
	mode := cipher.NewCBCDecrypter(block, c.InitVector)
	mode.CryptBlocks(decryptedMessage, encryptedMessage)
	return decryptedMessage, nil
}

// DecryptSimpleWithBase64 func
func (c *G_AES256Cryptor) DecryptSimpleWithBase64(base64Data string) ([]byte, error) {
	encryptedMessage, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	return c.Decrypt(encryptedMessage)
}

func (c *G_AES256Cryptor) DesDecryption(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.SecretKey)
	if err != nil {
		return nil, nil
	}

	blockMode := cipher.NewCBCDecrypter(block, c.InitVector)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// DesDecryptSimpleWithBase64 func
func (c *G_AES256Cryptor) DesDecryptSimpleWithBase64(base64Data string) ([]byte, error) {
	encryptedMessage, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	return c.DesDecryption(encryptedMessage)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
