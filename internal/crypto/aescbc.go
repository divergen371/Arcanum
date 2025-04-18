package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// PKCS#7 パディング
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func unpad(data []byte) ([]byte, error) {
	n := int(data[len(data)-1])
	if n > len(data) {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:len(data)-n], nil
}

// AES-CBC 暗号化
func EncryptCBC(key, iv, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	p := pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(p))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, p)
	return ciphertext, nil
}

// AES-CBC 復号化
func DecryptCBC(key, iv, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}
	plain := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plain, ciphertext)
	return unpad(plain)
}
