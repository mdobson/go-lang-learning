package hello

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func Hello() string {
	return "Hello!\n"
}

func Foo() string {
	return "Foo Response\n"
}

func RandToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x\n", b)
}

func GenerateKeyObjects() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		println(err.Error())
	}
	return privateKey
}

func GenerateKeys() string {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		println(err.Error())
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}

	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKeyDer, err := ssh.NewPublicKey(&privateKey.PublicKey)
	publicKeyPem := ssh.MarshalAuthorizedKey(publicKeyDer)
	return fmt.Sprintf("-----BEGIN RSA PUBLIC KEY-----\n%s-----END RSA PUBLIC KEY-----\n%s\n", string(publicKeyPem), privateKeyPem)
}
