# encrypt-and-decode-of-golang-and-node
golang 和 node 的加解密


AES-128-cbc加解密

    go使用方法：先运行go mod tidy;然后运行go run main.go
    node使用方法：node app.js

RSA加解密
    
    NodeJs版本RSA加密:
        命令: node.exe app.js RSA 0 text
        备注: text为待加密的明文字符串

    NodeJs版本RSA解密:
        命令: node.exe app.js RSA 1 text
        备注: text为待解密的密文字符串
    
    
    Go版本RSA加密:
        命令: go run main.go RSA 0 text
        备注: text为待加密的明文字符串

    Go版本RSA解密:
        命令: go run main.go RSA 1 text
        备注: text为待解密的密文字符串
