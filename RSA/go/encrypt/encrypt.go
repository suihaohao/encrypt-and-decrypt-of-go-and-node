package encrypt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func RsaEncrypt(data []byte) string {
	pubKey, err := ioutil.ReadFile("../public.pem")
	if err != nil {
		log.Fatal(err.Error())
	}
	block, _ := pem.Decode(pubKey) //将密钥解析成公钥实例
	if block == nil {
		log.Error("public key error")
		return ""
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		log.Error(err)
		return ""
	}
	pub := pubInterface.(*rsa.PublicKey)
	res, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, data, nil)
	if err != nil {
		log.Error(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(res)
}
