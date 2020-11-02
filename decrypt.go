package wxgameod

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

const pkcs7blocksize = 32

// Decrypt decrypt info
// Copy from  https://github.com/medivhzhan/miniapp
func Decrypt(sessionKey, encryptedData, iv string) ([]byte, error) {
	raw, err := decryptUserData(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// decryptUserData 解密用户数据
func decryptUserData(ssk, ciphertext, iv string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return nil, err
	}

	cipherStr, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	return cbcDecrypt(key, cipherStr, rawIV)
}

// cbcDecrypt CBC解密数据
func cbcDecrypt(key, ciphertext, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := aes.BlockSize
	iv = iv[:size]
	// ciphertext = ciphertext[size:] TODO: really useless?

	if len(ciphertext) < size {
		return nil, errors.New("ciphertext too short")
	}

	if len(ciphertext)%size != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7decode(ciphertext), nil
}

// pkcs7decode
func pkcs7decode(plaintext []byte) []byte {
	ln := len(plaintext)

	// 获取最后一个字符的 ASCII
	pad := int(plaintext[ln-1])
	if pad < 1 || pad > pkcs7blocksize {
		pad = 0
	}

	return plaintext[:(ln - pad)]
}
