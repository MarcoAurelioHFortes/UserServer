package auth

import (
	"MagicTableAPI/types"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

func GetPassword(user types.RegisterUserPayload, salt string) (string, error) {
	key, err := scrypt.Key([]byte(user.Password), []byte(salt), 1048576, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

func generateKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func GetSalt(user types.RegisterUserPayload) (string, error) {
	key, err := generateKey()
	if err != nil {
		return "", err
	}
	salt, err := encrypt(key, []byte(user.Email))
	if err != nil {
		return "", err
	}
	return string(salt), nil
}

func encrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}
