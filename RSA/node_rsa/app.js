const crypto = require('crypto');
const fs = require('fs');
const path = require('path');

const helpText = "参数错误,请采用以下方式运行程序\nnode.exe app.js RSA mode text\nmode: 0加密, 1解密\ntext: 明文或者密文\n示例 node.exe app.js RSA 0 abcdefg";

if (process.argv.length < 4) {
    console.log(helpText);
    return;
}

const method = process.argv[2];
const mode = parseInt(process.argv[3]);
if (method !== 'AES') {
    console.log(helpText);
    return;
}
if (mode !== 0 && mode !== 1) {
    console.log(helpText);
    return;
}

var text = '';
var password = '';

if (method === 'AES') {
    if (process.argv.length < 6) {
        console.log(helpText);
        return;
    }
    password = process.argv[4];
    text = process.argv[5];
} else {
    text = process.argv[4];
}

if (method === 'RSA' && mode === 0) {
    const publicKey = fs.readFileSync(path.join(__dirname, "../public.pem"), "utf8");
    const sign = crypto.publicEncrypt(publicKey, Buffer.from(text)).toString("base64");
    console.log("加密成功");
    console.log("明文:%s", text);
    console.log("密文:%s", sign);
} else if (method === 'RSA' && mode === 1) {
    var plainText = "";
    try {
        const privateKey = fs.readFileSync(path.join(__dirname, "../private.pem"), "utf8");
        plainText = crypto.privateDecrypt(privateKey, new Buffer(text, "base64")).toString("utf8");
    } catch (err) {
        console.log("解密失败, 密文:%s", text);
        return;
    }

    console.log("解密成功");
    console.log("密文:%s", text);
    console.log("明文:%s", plainText);
} 
