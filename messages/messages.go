package messages

import (
	"bytes"
	"chat-client/users"
	"chat-client/utils"
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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
	host := os.Getenv("CHAT_HOST")
	http.Post(host + "/messages", "application/json", bytes.NewBuffer(r))
}

func GetMessagePaload(messagePayload string) string{
	msg := utils.DecodeBase64String(messagePayload, *getPrivateKey())
	return msg
}

func getPrivateKey() *rsa.PrivateKey{
	b, _ := ioutil.ReadFile("user-data/privateKey.txt")
	return utils.GetPrivateKeyFromString(string(b))
}

func getDestinationPublicKey(destinationId string) *rsa.PublicKey{
	user := users.FindUserById(destinationId)
	return utils.GetPublicKeyFromString(user.PublicKey)
}

func GetUnreadMessages(userId string) []Message{
	host := os.Getenv("CHAT_HOST")
	r, err := http.Get( host + "/messages/unread/" + userId)
	if(err != nil){
		panic(err.Error())
	}
	var messages []Message
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&messages)
	return messages
}