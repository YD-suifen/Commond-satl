package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"strconv"
)

var listminio map[int]string
var xzlist []string

var url string = "https://172.16.204.246:8899"

type Login_token struct {
	Return []Eauth2 `json:"return"`
}

type Eauth2 struct {
	Eauth string `json:"eauth"`
	Expire float32 `json:"expire"`
	User string `json:"user"`
	Start string `json:"start"`
	Token string `json:"token"`
	Perms []string `json:"perms"`

}

type Api_login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth string `json:"eauth"`
}

func Token() (token string) {
	var login Api_login
	login.Username = "jiange"
	login.Password = "jiange123"
	login.Eauth = "pam"

	loginjson, err := json.Marshal(login)

	if err != nil {
		fmt.Println(err)
	}

	loginrequestinfo, err2 := http.NewRequest("POST","https://172.16.204.246:8899/login", bytes.NewBuffer(loginjson))
	if err2 != nil {
		fmt.Println(err2)
	}
	loginrequestinfo.Header.Set("Accept", "application/json")
	loginrequestinfo.Header.Set("Content-Type", "application/json")


	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
	}
	respon, err3 := client.Do(loginrequestinfo)

	if err3 != nil {
		fmt.Println(err3)
	}

	body, err4 := ioutil.ReadAll(respon.Body)
	if err4 != nil {
		fmt.Println(err4)
	}

	var token2 Login_token
	var nowtoken string
	err5 := json.Unmarshal(body, &token2)
	if err5 != nil {
		fmt.Println(err5)
	}

	for _, v := range token2.Return {
		nowtoken = v.Token
	}

	return nowtoken

}

type KeylistRespon struct {

	Return []struct{
		Tag string `json:"tag"`
		Data struct{
			Jid string `json:"jid"`
			Return struct{
				Local []string `json:"local"`
				Minions_rejected []interface{} `json:"minions_rejected"`
				Minions_denied []interface{} `json:"minions_denied"`
				Minions_pre []interface{} `json:"minions_pre"`
				Minions []string `json:"minions"`
			} `json:"return"`
			Success bool `json:"success"`
			Stamp string `json:"_stamp"`
			Tag string `json:"tag"`
			User string `json:"user"`
			Fun string `json:"fun"`

		} `json:"data"`
	} `json:"return"`
}


type Request struct {
	Client string `json:"client"`
	Tgt string `json:"tgt"`
	Fun string `json:"fun"`
	Arg string `json:"arg"`
	Match string `json:"match"`
}

func RequestKeylist() []string {
	var request Request
	request.Client = "wheel"
	request.Tgt = "*"
	request.Fun = "key.list_all"

	requesjson, err := json.Marshal(request)

	if err != nil {
		fmt.Println(err)
	}
	httreques, err := http.NewRequest("POST",url,bytes.NewBuffer(requesjson))
	if err !=nil {
		fmt.Println(err)
	}

	tocken := Token()
	httreques.Header.Set("Accept", "application/json")
	httreques.Header.Set("Content-Type", "application/json")
	httreques.Header.Set("X-Auth-Token", tocken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	respone, err := client.Do(httreques)
	if err !=nil{
		fmt.Println(err)
	}
	var Respon KeylistRespon
	byterespon, err := ioutil.ReadAll(respone.Body)
	if err !=nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(byterespon,&Respon)

	var keylist []string

	for _, v := range Respon.Return {
		keylist = v.Data.Return.Minions
	}
	return keylist

}


type CPCommod struct {
	Client string `json:"client"`
	Tgt []string `json:"tgt"`
	Fun string `json:"fun"`
	Arg []string `json:"arg"`
	Match string `json:"match"`
}

type ActionRepone struct {
	Return []map[string]string `json:"return"`
}

func CPDir(Targethost []string, srcname string, targetpath string)  bool {

	var request CPCommod
	request.Fun = "cp.get_dir"
	request.Tgt = Targethost
	request.Client = "local"
	request.Arg = append(request.Arg,"salt://" + srcname)
	request.Arg = append(request.Arg, targetpath)

	jsonreques, err := json.Marshal(request)
	if err !=nil{
		fmt.Println(err)
	}
	Request, err := http.NewRequest("POST",url,bytes.NewBuffer(jsonreques))
	if err != nil {
		fmt.Println(err)
	}
	tocken := Token()
	Request.Header.Set("Accept", "application/json")
	Request.Header.Set("Content-Type", "application/json")
	Request.Header.Set("X-Auth-Token", tocken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}

	Respon, err := client.Do(Request)
	if err !=nil{
		fmt.Println(err)
	}
	var respone ActionRepone
	respones,err := ioutil.ReadAll(Respon.Body)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(respones, &respone)

	fmt.Println(respone.Return)
	return true

}






func main()  {

	var actionhostlist []string

	fmt.Println("以下是列表，选择,逗号相隔，或者选择ALL")
	list := RequestKeylist()
	for i, v :=range list{
		fmt.Printf("ID: %d,--Server: %v", i, v)
	}

	fmt.Println(">input host")

	f := bufio.NewReader(os.Stdin)


	input, _ := f.ReadString('\n')

	fmt.Println(">input srcname")

	b := bufio.NewReader(os.Stdin)


	inputsrcname, _ := b.ReadString('\n')

	fmt.Println(">input dirpath")

	c := bufio.NewReader(os.Stdin)


	inputdirpath, _ := c.ReadString('\n')

	if input == "all" {


	}else {
		xzlist = strings.Split(input,",")

		fmt.Println("list",xzlist)

		fmt.Println(input)

		for _, v := range xzlist{
			v,_  := strconv.Atoi(v)
			actionhostlist = append(actionhostlist, list[v])
		}



		what := CPDir(actionhostlist, inputsrcname, inputdirpath)

		if what {
			fmt.Println("true  true")
			return
		}else {
			fmt.Println("fales")
			return
		}


	}





}



//curl -k https://172.16.204.246:8899 -H "Accept: application/x-yaml" -H "X-Auth-Token: 8d95a58dfb90594a4f63187d4a84d9b97826d281" -d client='local' -d tgt='*' -d fun='file.copy' -d arg='/srv/salt/test' -d arg='/tmp/test' -d arg='recurse=True'
//
//
//curl -k https://172.16.204.246:8899 -H "Accept: application/x-yaml" -H "X-Auth-Token:db67d6a974dd80f03038df018e757371b145f43a" -d client='local' -d tgt='*' -d fun='cp.get_dir' -d arg='salt://test' -d arg='dest=/tmp/'







































































