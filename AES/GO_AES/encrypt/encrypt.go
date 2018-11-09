package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"time"
	log "github.com/sirupsen/logrus"
)
var commonStr = "1234567890~!@#$%^*abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func AesEncrypt(password string, data interface{}) string {
	key, iv := ByteToKey(password, 16)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Error: NewCipher(%d bytes) = %s", len(password), err)
	}
	dataByte, _ := json.Marshal(data)

	content := PKCS5Padding(dataByte, block.BlockSize())
	enc := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(content))
	enc.CryptBlocks(crypted, content)

	return base64.StdEncoding.EncodeToString(crypted)
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func ByteToKey(password string, keylen int) ([]byte, []byte) {
	pass := []byte(password)
	prev := []byte{}
	key := []byte{}
	iv := []byte{}

	remain := 0
	for len(key) < keylen {
		hash := md5.Sum(append(prev, pass...))
		remain = keylen - len(key)
		if remain < 16 {
			key = append(key, hash[:remain]...)
		} else {
			key = append(key, hash[:]...)
		}
		prev = hash[:]
	}

	hash := md5.Sum(append(prev, pass...))
	if remain < 16 {
		iv = append(prev[remain:], hash[:remain]...)
	} else {
		iv = hash[:]
	}

	return key, iv
}

func MakePwd(pwdl int) (string) {
	rand.Seed(time.Now().UnixNano())
	commonByte := []byte(commonStr)
	str := []byte{}
	for i := 0; i < pwdl; i++ {
		num := rand.Intn(len(commonStr))
		str = append(str, commonByte[num])
	}
	return string(str)
}
