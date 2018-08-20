package http

import (
	"net/http"
	"bytes"
	"fmt"
	"crypto/tls"
	"io/ioutil"
)

var Url string = "https://172.16.204.246:8899"
var Loginurl string  = "https://172.16.204.246:8899/login"


func Httprequest(action, url string, body []byte, token ...string) (error, []byte) {

	request, err := http.NewRequest(action,url,bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	if len(token) > 0 {
		request.Header.Set("X-Auth-Token",token[0])
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify:true,
		},
	}
	client :=&http.Client{
		Transport:tr,
	}
	respones, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	sponebody, _ := ioutil.ReadAll(respones.Body)

	return nil, sponebody

}
