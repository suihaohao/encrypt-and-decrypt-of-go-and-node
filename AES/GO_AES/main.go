package main

import (
	log "github.com/sirupsen/logrus"
	"go_aes/decrypt"
	"go_aes/encrypt"
)

type TestData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Content  string `json:"content"`
}

func main() {
  data := TestData{}
  data.Name = "suihaohao"
  data.Password = "1234565"
  data.Content = "梅须逊雪三分白,雪却输梅一段香"
  password := encrypt.makePwd(16)
  encData := encrypt.AesEncrypt(password, data)
  
  log.println("encData:", encData)
  
  decData := decrypt(base64.StdEncoding.DecodeString(encData), []byte(password))
  
  log.println("decData:", decData)
  
}
