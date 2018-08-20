package token

import (
	"encoding/json"
	"fmt"
	"Commond-satl/http"
)

func Token() (string, error) {
	var login Api_login
	login.Username = "jiange"
	login.Password = "jiange123"
	login.Eauth = "pam"

	loginjson, _ := json.Marshal(login)

	err, sponebody := http.Httprequest("POST",http.Loginurl, loginjson)
	if err !=nil {
		fmt.Println("token repuest err:",err)
		return "", err
	}

	var token2 Login_token
	var nowtoken string
	err5 := json.Unmarshal(sponebody, &token2)
	if err5 != nil {
		fmt.Println("token json unmarshal err:",err5)
		return "", err5
	}

	for _, v := range token2.Return {
		nowtoken = v.Token
	}

	return nowtoken, nil

}
