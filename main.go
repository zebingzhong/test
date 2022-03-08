package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type R struct {
	Encrypt string
}

type S struct {
	Challenge string
	Token     string
	Type      string
}

func main() {
	r := gin.Default()

	r.POST("/callback", func(c *gin.Context) {
		//buf := make([]byte, 1024)
		//n, _ := c.Request.Body.Read(buf)
		//s := string(buf[0:n])
		json := R{}
		err := c.BindJSON(&json)
		if err != nil {
			return
		}
		fmt.Println(json)
		if err != nil {
			panic("解析错误")
		}
		s, err := Decrypt(json.Encrypt, "5B6Bak3DSat4QAOY4VZSWb07f3I0RXEK")
		if err != nil {
			panic(err)
		}
		rJson := S{}
		if err = json2.Unmarshal([]byte(s), &rJson); err != nil {
			fmt.Println(rJson)
		}
		fmt.Println(rJson)
		// 解析
		c.JSON(200, gin.H{
			"challenge": rJson.Challenge,
		})
	})
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Decrypt(encrypt string, key string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := strings.Index(string(buf), "{")
	if n == -1 {
		n = 0
	}
	m := strings.LastIndex(string(buf), "}")
	if m == -1 {
		m = len(buf) - 1
	}
	return string(buf[n : m+1]), nil
}
