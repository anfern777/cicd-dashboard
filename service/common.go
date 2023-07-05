package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDB(c *gin.Context) *gorm.DB {
	return c.Keys["DB"].(*gorm.DB)
}

func encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	b := []byte(text)
	b = PKCS5Padding(b, block.BlockSize())
	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(b))
	mode.CryptBlocks(ciphertext, b)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedtext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	if len(decodedtext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext block size is too short")
	}

	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(decodedtext, decodedtext)

	return string(PKCS5UnPadding(decodedtext)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
