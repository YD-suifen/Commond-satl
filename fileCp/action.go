package fileCp

import (
	"encoding/json"
	"fmt"
	"Commond-satl/http"
	"Commond-satl/token"

)

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
	tocken := token.Token()
	err, respones := http.Httprequest("POST",http.Url,jsonreques,tocken)
	if err != nil {
		fmt.Println(err)
	}

	var respone saltstack.ActionRepone
	_ = json.Unmarshal(respones, &respone)

	fmt.Println(respone.Return)
	return true

}