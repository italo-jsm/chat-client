package tests

import (
	"chat-client/utils"
	"testing"
)

func TestEncryptAndDecriptString(t *testing.T) {
	publicKey, privateKey := utils.GenerateKeyPair()
	stringTest := "Palavras de Teste"
	encrypted := utils.EncryptString(stringTest, *utils.GetPublicKeyFromString(publicKey))
	decrypted := utils.DecodeBase64String(encrypted, *utils.GetPrivateKeyFromString(privateKey))
	if(stringTest != decrypted){
		t.Fail()
	}
}

func TestEncryptReturnsDiferentResults(t *testing.T) {
	publicKey, _ := utils.GenerateKeyPair()
	stringTest := "Palavras de Teste"
	encrypted1 := utils.EncryptString(stringTest, *utils.GetPublicKeyFromString(publicKey))
	encrypted2 := utils.EncryptString(stringTest, *utils.GetPublicKeyFromString(publicKey))
	if(encrypted1 == encrypted2){
		t.Fail()
	}
}