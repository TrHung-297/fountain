package crypto

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestG_AES256_Crypt(t *testing.T) {
	aes := newG_AES256Cryptor([]byte("scPQ4l2lHNZpWQwn8vgH0QUr4iQw9mHr"), []byte("jKPOOoiAW2MfNxRN"))

	message := fmt.Sprintf("This is test %d", rand.Int63())
	encryptedMessage, err := aes.Encrypt([]byte(message))
	assert.Equal(t, err, nil)
	assert.Equal(t, len(message), len(encryptedMessage))

	msg, err := aes.Decrypt(encryptedMessage)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(encryptedMessage), len(msg))
	assert.Equal(t, message, string(msg))
}
