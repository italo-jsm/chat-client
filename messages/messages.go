package messages

import (
	"bytes"
	"chat-client/users"
	"chat-client/utils"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendMessage(destinationId string, messagePayload string){
	//pegar chave privada do destinatario
	publicKey := getDestinationPublicKey(destinationId)
	//encriptar mensagem
	encryptedMessage := utils.EncryptString(messagePayload, *publicKey)
	//fazer post para o servico de mensagens
	message := Message{
		SenderId: "goApp",
		ReceiverId: destinationId,
		Payload: encryptedMessage,
	}
	r, _ := json.Marshal(message)
	http.Post("http://localhost:3000/messages", "application/json", bytes.NewBuffer(r))
}

func ReceiveMessage(messagePayload string){
	msg := utils.DecodeBase64String(messagePayload, *getPrivateKey())
	println(msg)
}

func getPrivateKey() *rsa.PrivateKey{
	//buscar chave privada no storage local
	b, _ := ioutil.ReadFile("user-data/privateKey.txt")
	//montar chave privada
	bytes, err := base64.StdEncoding.DecodeString(string(b))
	if(err != nil){
		panic(err.Error())
	}
	privateKey, err2 := x509.ParsePKCS1PrivateKey(bytes)
	if(err2 != nil){
		panic(err2.Error())
	}
	return privateKey
}

func getDestinationPublicKey(destinationId string) *rsa.PublicKey{
	//chamar o servico de mensagens e obter o usuario
	user := users.FindUserById(destinationId)
	//montar chave publica
	bytes, err := base64.StdEncoding.DecodeString(user.PublicKey)
	if(err != nil){
		panic(err.Error())
	}
	publicKey, err2 := x509.ParsePKCS1PublicKey(bytes)
	if(err2 != nil){
		panic(err2.Error())
	}
	return publicKey
}