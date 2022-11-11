package adapter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

// Encryptor encrypts/decrypts data.
type Encryptor struct {
	CipherPass string
}

func NewEncryptor(cipherPass string) *Encryptor {
	return &Encryptor{CipherPass: cipherPass}
}

func (e *Encryptor) Encrypt(data []byte) (string, error) {
	aesgcm, nonce, err := e.getAesgcm()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(aesgcm.Seal(nil, nonce, data, nil)), nil
}

func (e *Encryptor) Decrypt(encrypted string) ([]byte, error) {
	aesgcm, nonce, err := e.getAesgcm()
	if err != nil {
		return nil, err
	}

	decoded, err := hex.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	decrypted, err := aesgcm.Open(nil, nonce, decoded, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func (e *Encryptor) getAesgcm() (cipher.AEAD, []byte, error) {
	key := sha256.Sum256([]byte(e.CipherPass))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, nil, err
	}

	nonce := key[len(key)-aesgcm.NonceSize():]

	return aesgcm, nonce, nil
}
