package token


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


type Login_token struct {
	Return []Eauth2 `json:"return"`
}