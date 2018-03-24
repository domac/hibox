package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path"
)

func genRSAKey(bits int) ([]byte, []byte) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	private_key := x509.MarshalPKCS1PrivateKey(key)
	public_key, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)

	//生成私钥
	bprivate := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: private_key,
	}

	//生成公钥
	bpublic := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: public_key,
	}

	private_pem := pem.EncodeToMemory(bprivate)
	public_pem := pem.EncodeToMemory(bpublic)
	return private_pem, public_pem
}

//处理文件
func handleFile(from, to string) error {

	file, err := os.Open(from)
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil
	}

	size := fileInfo.Size()
	println(size)

	return nil
}

func main() {
	private_pem, public_pem := genRSAKey(1024)
	fmt.Printf("private pem: %s\npublic pem:%s", private_pem, public_pem)

	wd, _ := os.Getwd()
	sourceFilePath := path.Join(wd, "test.txt")
	targetFilePath := path.Join(wd, "test2.txt")

	err := handleFile(sourceFilePath, targetFilePath)
	if err != nil {
		fmt.Println(err.Error())
	}
}
