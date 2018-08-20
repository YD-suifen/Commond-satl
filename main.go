package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"Commond-satl/saltnode"
	"Commond-satl/fileCp"
)

var listminio map[int]string
var xzlist []string


func main()  {

	var actionhostlist []string

	fmt.Println("以下是列表，选择,逗号相隔，或者选择ALL")
	list := saltnode.RequestKeylist()
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



		what := fileCp.CPDir(actionhostlist, inputsrcname, inputdirpath)

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







































































