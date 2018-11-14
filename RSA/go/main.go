package main

import (
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"go-crypto/decrypt"
	"go-crypto/encrypt"
	"os"
)

const helpText = "参数错误,请采用以下方式运行程序\ngo run main.go RSA mode text\nmode: 0加密, 1解密\ntext: 明文或者密文\n示例 go run main.go RSA 0 abcdefgg"

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	if len(os.Args) == 4 {
		method := os.Args[1]
		mod := os.Args[2]
		text := os.Args[3]
		if method != "RSA" {
			log.Println(helpText)
			return
		}
		if mod == "0" {
			encrypted := encrypt.RsaEncrypt([]byte(text))
			log.Println("加密成功")
			log.Println("明文:", text)
			log.Println("密文:", encrypted)
			return
		}
		if mod == "1" {
			decrypted, err := decrypt.RsaDecrypt(text)
			if err != nil {
				log.Println("解密失败")
				return
			}
			log.Println("解密成功")
			log.Println("密文:", text)
			log.Println("明文:", decrypted)
			return
		}
		log.Println(helpText)
	}
	log.Println(helpText)
}
