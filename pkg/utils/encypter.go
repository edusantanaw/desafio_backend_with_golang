package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const saltSize = 16
const keySize = 32
const iterations = 100000

func deriveKey(secret string, salt []byte) []byte {
	return pbkdf2.Key([]byte(secret), salt, iterations, keySize, sha256.New)
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func Encrypt(plainText, secret string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	key := deriveKey(secret, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	finalData := append(salt, cipherText...) // Salt + CipherText
	return base64.StdEncoding.EncodeToString(finalData), nil
}

func Decrypt(encryptedText, secret string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	if len(data) < saltSize {
		return "", errors.New("invalid encrypted text")
	}

	salt, cipherData := data[:saltSize], data[saltSize:]
	key := deriveKey(secret, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherData) < nonceSize {
		return "", errors.New("invalid encrypted text")
	}

	nonce, cipherText := cipherData[:nonceSize], cipherData[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
