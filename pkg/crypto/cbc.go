package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

const key = "1234567890qwertyuio!@#$%^&*()-+?"

// Encrypt make original string encrypted
func Encrypt(orig string) (string, error) {
	k := []byte(key)
	origByte := []byte(orig)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", fmt.Errorf("key length must in 16/24/32: %v", err)
	}

	blockSize := block.BlockSize()
	origByte = pkcs7Pad(origByte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, biv(blockSize))
	cryted := make([]byte, len(origByte))
	blockMode.CryptBlocks(cryted, origByte)
	return base64.StdEncoding.EncodeToString(cryted), nil
}

// Decrypt make encrypted string into original
func Decrypt(cryted string) (string, error) {
	k := []byte(key)
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return "", fmt.Errorf("invalid cryted: %v", err)
	}

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", fmt.Errorf("key length must in 16/24/32: %v", err)
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, biv(blockSize))
	orig := make([]byte, len(crytedByte))
	blockMode.CryptBlocks(orig, crytedByte)
	orig = pkcs7Unpad(orig)
	return string(orig), nil
}

// Decryptniv make encrypted string into original without biv
func Decryptniv(cryted string) (string, error) {
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return "", fmt.Errorf("invalid cryted: %v", err)
	}

	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", fmt.Errorf("key length must in 16/24/32: %v", err)
	}

	size := block.BlockSize()
	mode := cipher.NewCBCDecrypter(block, k[:size])
	orig := make([]byte, len(crytedByte))
	mode.CryptBlocks(orig, crytedByte)
	orig = pkcs7Unpad(orig)
	return string(orig), nil
}

func pkcs7Pad(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7Unpad(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]

}

func biv(size int) []byte {
	md := md5.New()
	md.Write([]byte(time.Now().Format("2006/01/02")))
	hash := hex.EncodeToString(md.Sum(nil))
	return []byte(hash)[:size]
}
