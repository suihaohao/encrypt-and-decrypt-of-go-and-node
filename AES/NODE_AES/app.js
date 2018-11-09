const crypto = require('crypto');
const text = JSON.stringify({
        "name": "my name",
        "password": "1234565",
        "content": "梅须逊雪三分白,雪却输梅一段香"
    })
function encrypt_aes(key, message) {
    const cipher = crypto.createCipher('aes-128-cbc', key);
    var sign = cipher.update(message, 'utf8', 'base64');
    sign += cipher.final('base64');
    console.log("加密成功");
    return sign
}

function decrypt_aes(key, message) {
    var plainText = ""
    const decipher = crypto.createDecipher('aes-128-cbc', key);
    plainText = decipher.update(message, 'base64', 'utf8');
    plainText += decipher.final('utf8');
    console.log("解密成功");
    return plainText
}

function test(){
  const password = "123456789"
  encData = encrypt_aes(password, text)
  console.log("encData:", encData)
  decData = decrypt_aes(password, encData)
  console.log("decData:", decData)
}

test()
