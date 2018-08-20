package saltnode

import (
	"encoding/json"
	"fmt"
	"Commond-satl/saltstack"
	"Commond-satl/saltstack/token"
	"Commond-satl/saltstack/http"
)

func RequestKeylist() []string {
	var request saltstack.Request
	request.Client = "wheel"
	request.Tgt = "*"
	request.Fun = "key.list_all"

	requesjson, err := json.Marshal(request)

	if err != nil {
		fmt.Println(err)
	}

	tocken, err := token.Token()

	err, respone := http.Httprequest("POST", http.Url, requesjson, tocken)
	if err != nil {
		fmt.Println(err)
	}

	var Respon saltstack.KeylistRespon

	_ = json.Unmarshal(respone,&Respon)

	var keylist []string

	for _, v := range Respon.Return {
		keylist = v.Data.Return.Minions
	}
	return keylist

}