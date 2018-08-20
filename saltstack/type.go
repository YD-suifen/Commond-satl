package saltstack


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

type ActionRepone struct {
	Return []map[string]string `json:"return"`
}




