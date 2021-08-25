// Code generated by goctl. DO NOT EDIT.
package types

type CreateUserReq struct {
	UserName string `json:"userName"`
	Pwd      string `json:"pwd"`
}

type CreateUserResp struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    OneUser `json:"data"`
}

type FindUserResp struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    OneUser `json:"data"`
}

type OneUser struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
