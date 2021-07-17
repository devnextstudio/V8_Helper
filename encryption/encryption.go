package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
)

type Crypto interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherIvKey string) (string, error)
}

type niceCrypto struct {
	cipherKey   string
	cipherIvKey string
}

var crypt Crypto
var once sync.Once

var cipherKey = "CIPHERKEY01234567890123456789012"
var cipherIvKey = "CIPHERIVKEY01234"

func GetCrypto() Crypto {
	once.Do(func() {
		c, err := NewNiceCrypto(cipherKey, cipherIvKey)
		if err != nil {
			fmt.Println("NewNiceCrypto: ")
			fmt.Println(err)
			return
		}
		crypt = c
	})
	return crypt
}

func (c niceCrypto) Encrypt(plainText string) (string, error) {
	if strings.TrimSpace(plainText) == "" {
		return plainText, nil
	}

	block, err := aes.NewCipher([]byte(c.cipherKey))
	if err != nil {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(block, []byte(c.cipherIvKey))
	paddedPlainText := padPKCS7([]byte(plainText), encrypter.BlockSize())

	cipherText := make([]byte, len(paddedPlainText))
	// CryptBlocks 함수에 데이터(paddedPlainText)와 암호화 될 데이터를 저장할 슬라이스(cipherText)를 넣으면 암호화가 된다.
	encrypter.CryptBlocks(cipherText, paddedPlainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (c niceCrypto) Decrypt(cipherText string) (string, error) {
	if strings.TrimSpace(cipherText) == "" {
		return cipherText, nil
	}

	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(c.cipherKey))
	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCBCDecrypter(block, []byte(c.cipherIvKey))
	plainText := make([]byte, len(decodedCipherText))

	decrypter.CryptBlocks(plainText, decodedCipherText)
	trimmedPlainText := trimPKCS5(plainText)

	return string(trimmedPlainText), nil
}

func NewNiceCrypto(cipherKey string, cipherIvKey string) (Crypto, error) {
	if ck := len(cipherKey); ck != 32 {
		return nil, aes.KeySizeError(ck)
	}

	if cik := len(cipherIvKey); cik != 16 {
		return nil, aes.KeySizeError(cik)
	}

	crypt = &niceCrypto{cipherKey, cipherIvKey}

	return crypt, nil
}

func padPKCS7(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padText...)
}

func trimPKCS5(text []byte) []byte {
	padding := text[len(text)-1]
	return text[:len(text)-int(padding)]
}
