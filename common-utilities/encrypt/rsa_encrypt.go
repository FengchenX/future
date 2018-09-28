//author xinbing
//time 2018/8/25 15:31
//rsa工具
package encrypt

import (
	"encoding/pem"
	"github.com/pkg/errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
)
// Encrypt
func RsaEncrypt(content string,publicKey string) ([]byte,error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	contentBytes := []byte(content)
	resultBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, contentBytes)
	return resultBytes, err
}

// Decrypt
func RsaDecrypt(content string,privateKey string) ([]byte,error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	privateInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	contentBytes := []byte(content)
	resultBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateInterface, contentBytes)
	return resultBytes,err
}

// Generate RSA PrivateKey PublicKey
func GenKeyPairs(bits int) (privateKey ,publicKey string,err error) {
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	block := &pem.Block{
		Type:  "private key",
		Bytes: derStream,
	}
	b := pem.EncodeToMemory(block)
	privateKey = string(b)

	pubKey := &priKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "public key",
		Bytes: derPkix,
	}
	b = pem.EncodeToMemory(block)
	publicKey = string(b)
	return privateKey, publicKey, nil
}