package users

import (
	"bytes"
	"chat-client/utils"
	"encoding/json"
	"net/http"
	"os"
)

func CreateUser(username string, email string){
	publicKey, privateKey := utils.GenerateKeyPair()
	os.WriteFile("user-data/privateKey.txt", []byte(privateKey), 0755)
	user := User{
		Username: username,
		Email: email,
		PublicKey: publicKey,
	}
	r, _ := json.Marshal(user)
	host := os.Getenv("CHAT_HOST")
	http.Post( host + "/users", "application/json", bytes.NewBuffer(r))
}

func FindUserById(userId string) User{
	host := os.Getenv("CHAT_HOST")
	r, err := http.Get( host + "/users/" + userId)
	if(err != nil){
		panic(err.Error())
	}
	user := User{}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&user)
	return user
}