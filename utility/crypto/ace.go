package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type uAes struct{}

var Aes = uAes{}

// pkcs7Padding 使用PKCS7进行填充，IOS也是7
func (u *uAes) pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesCBCEncrypt aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func (u *uAes) AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建AES cipher失败: %w", err)
	}

	// 填充原文
	blockSize := block.BlockSize()
	rawData = u.pkcs7Padding(rawData, blockSize)
	// 初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	// block大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("生成随机IV失败: %w", err)
	}

	// block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func (u *uAes) AesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建AES cipher失败: %w", err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		return nil, fmt.Errorf("密文长度不足，至少需要%d字节，实际%d字节", blockSize, len(encryptData))
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		return nil, fmt.Errorf("密文长度必须是块大小的倍数，块大小%d，密文长度%d", blockSize, len(encryptData))
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = pkcs7UnPadding(encryptData)
	return encryptData, nil
}

func (u *uAes) Encrypt(rawData, key []byte) (string, error) {
	data, err := u.AesCBCEncrypt(rawData, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func (u *uAes) Decrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := u.AesCBCDecrypt(data, key)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// Dncrypt 保持向后兼容（已废弃，请使用Decrypt）
func (u *uAes) Dncrypt(rawData string, key []byte) (string, error) {
	return u.Decrypt(rawData, key)
}

// AesCBCDncrypt 保持向后兼容（已废弃，请使用AesCBCDecrypt）
func (u *uAes) AesCBCDncrypt(encryptData, key []byte) ([]byte, error) {
	return u.AesCBCDecrypt(encryptData, key)
}
