package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"strconv"
)

type AesCTR128KeySizeError int

func (k AesCTR128KeySizeError) Error() string {
	return "AesCTR128KeySizeError: invalid key size " + strconv.Itoa(int(k))
}

type AesCTR128Encrypt struct {
	// block cipher.Block
	stream cipher.Stream
}

// NewAesCTR128Encrypt func;
// key length must be 16, 24 or 32
func NewAesCTR128Encrypt(key []byte, iv []byte) (*AesCTR128Encrypt, error) {
	block2, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// iv length must be 16
	if len(iv) != 16 {
		return nil, AesCTR128KeySizeError(len(iv))
	}

	stream2 := cipher.NewCTR(block2, iv)

	return &AesCTR128Encrypt{
		// block:	block2,
		stream: stream2,
	}, nil
}

func (aes *AesCTR128Encrypt) Encrypt(plaintext []byte) []byte {
	aes.stream.XORKeyStream(plaintext, plaintext)
	return plaintext
}
