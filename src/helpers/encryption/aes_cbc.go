package encryptionHelper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func encryptAESCBC(key []byte, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// generate random iv
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// padding plaintext to multiple of block size
	paddedPlaintext := padPlaintext(plaintext, aes.BlockSize)

	// create cipher with CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)

	// encrypt the plaintext
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	// append iv to the beginning of the ciphertext
	ciphertext = append(iv, ciphertext...)

	// encode the ciphertext to hex string
	ciphertextHex := hex.EncodeToString(ciphertext)

	return ciphertextHex, nil
}

func decryptAESCBC(key []byte, ciphertextHex string) (decrypted []byte, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New(rec.(string))
		}
	}()

	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// get iv from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// create cipher with CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)

	// decrypt the ciphertext
	paddedPlaintext := make([]byte, len(ciphertext))

	defer mode.CryptBlocks(paddedPlaintext, ciphertext)

	// unpad the plaintext
	plaintext, err := unpadPlaintext(paddedPlaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func padPlaintext(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func unpadPlaintext(paddedPlaintext []byte) ([]byte, error) {
	length := len(paddedPlaintext)
	padding := int(paddedPlaintext[length-1])
	if padding > length {
		return nil, errors.New("invalid padding")
	}
	return paddedPlaintext[:length-padding], nil
}
