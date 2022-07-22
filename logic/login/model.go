package login

type RegisterReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRsp struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginRsp struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}
