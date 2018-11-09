package main

import (
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"go-aes/decrypt"
	"go-aes/encrypt"
)

type TestData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Content  string `json:"content"`
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	data := TestData{}
	data.Name = "my name"
	data.Password = "1234565"
	data.Content = "梅须逊雪三分白,雪却输梅一段香"
	password := encrypt.MakePwd(16)
	encData := encrypt.AesEncrypt(password, data)

	log.Println("encData:", encData)
	message, _ := base64.StdEncoding.DecodeString(encData)
	decData := decrypt.AesDecrypt(message, []byte(password))
	log.Println("decData:", decData)

}
