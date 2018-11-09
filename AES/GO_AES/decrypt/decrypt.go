package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"go-aes/encrypt"

	log "github.com/sirupsen/logrus"
)

func AesDecrypt(ciphertext []byte, password []byte) (interface{}) {
	key, iv := encrypt.ByteToKey(string(password), 16)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Error: NewCipher(%d bytes) = %s", len(password), err)
	}
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	dec.CryptBlocks(decrypted, ciphertext)
	decrypted = PKCS5UnPadding(decrypted)
	var data map[string]interface{}
	err = json.Unmarshal(decrypted, &data)
	if err != nil {
		return string(decrypted)
	}
	return data
}

func PKCS5UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
