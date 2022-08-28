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
	os.WriteFile("privateKey.txt", []byte(privateKey), 0755)
	user := User{
		Username: username,
		Email: email,
		PublicKey: publicKey,
	}
	r, _ := json.Marshal(user)
	http.Post("http://localhost:3000/users", "application/json", bytes.NewBuffer(r))
}

func FindUserById(userId string) User{
	r, err := http.Get("http://localhost:3000/users/" + userId)
	if(err != nil){
		panic(err.Error())
	}
	user := User{}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&user)
	return user
}