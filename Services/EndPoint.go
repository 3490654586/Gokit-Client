package Services

type UserRequest struct {
	//封装User请求结构体
	Uid int `json:"uid"`
}

type UserResponse struct {
	Resule string `json:"resule"`
}
