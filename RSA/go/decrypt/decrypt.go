package decrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	log "github.com/sirupsen/logrus"
	"go-crypto/encrypt"
	"io/ioutil"
)

func RsaDecrypt(data string) (string, error) {
	priKey, err := ioutil.ReadFile("../private.pem")
	if err != nil {
		log.Fatal(err.Error())
	}
	block, _ := pem.Decode(priKey)
	if block == nil {
		return "", errors.New("private key error!")
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Error(err)
		return "", err
	}
	priv := pri.(*rsa.PrivateKey)
	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Error(err)
	}
	res, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, cipherText, nil)
	return string(res), err
}
