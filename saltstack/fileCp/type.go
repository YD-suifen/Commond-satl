package fileCp



type CPCommod struct {
	Client string `json:"client"`
	Tgt []string `json:"tgt"`
	Fun string `json:"fun"`
	Arg []string `json:"arg"`
	Match string `json:"match"`
}