/* !!
 * File: rsa_cryptor.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:30:40 am
 
 */

package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
)

// RSACryptor func;
type RSACryptor struct {
	key *rsa.PrivateKey
}

// NewRSACryptor func;
func NewRSACryptor(pkcs1PemPrivateKey []byte) *RSACryptor {
	block, _ := pem.Decode(pkcs1PemPrivateKey)
	if block == nil {
		err := fmt.Errorf("invalid pem key data")
		panic(err)
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err := fmt.Errorf("failed to parse private key: %s", err.Error())
		panic(err)
	}

	return &RSACryptor{
		key: key,
	}
}

// Encrypt func;
func (m *RSACryptor) Encrypt(b []byte) []byte {
	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(b), big.NewInt(int64(m.key.E)), m.key.N)

	e := c.Bytes()
	var size = len(e)
	if size < 256 {
		size = 256
	}

	res := make([]byte, size)
	copy(res, c.Bytes())

	return res
}

// Decrypt func;
func (m *RSACryptor) Decrypt(b []byte) []byte {
	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(b), m.key.D, m.key.N)
	return c.Bytes()
}
