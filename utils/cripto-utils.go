package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

func EncryptString(message string, key rsa.PublicKey) string{
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&key,
		[]byte(message),
		nil)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes)
}

func DecodeBase64String(base64String string, privateKey rsa.PrivateKey) string{
	result, _ := base64.StdEncoding.DecodeString(base64String)
	decryptedBytes, err := privateKey.Decrypt(nil, result, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	return string(decryptedBytes)
}

func GenerateKeyPair() (string, string){
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if(err != nil){
		panic(err.Error())
	}
	pubString := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	privString := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))

	return pubString, privString
}

func GetPublicKeyFromString(publicKeyString string) *rsa.PublicKey{
	bytes, err := base64.StdEncoding.DecodeString(publicKeyString)
	if(err != nil){
		panic(err.Error())
	}
	publicKey, err2 := x509.ParsePKCS1PublicKey(bytes)
	if(err2 != nil){
		panic(err2.Error())
	}
	return publicKey
}

func GetPrivateKeyFromString(stringPrivateKey string) *rsa.PrivateKey{
	bytes, err := base64.StdEncoding.DecodeString(stringPrivateKey)
	if(err != nil){
		panic(err.Error())
	}
	privateKey, err2 := x509.ParsePKCS1PrivateKey(bytes)
	if(err2 != nil){
		panic(err2.Error())
	}
	return privateKey
}
