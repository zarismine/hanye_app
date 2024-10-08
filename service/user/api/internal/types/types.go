// Code generated by goctl. DO NOT EDIT.
package types

type Default struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Login struct {
	Id     string `json:"id"`
	OpenId string `json:"openid"`
	Token  string `json:"token"`
}

type LoginReq struct {
	Code string `json:"code"`
}

type LoginResp struct {
	Default
	Data struct {
		Id     string `json:"id"`
		OpenId string `json:"openid"`
		Token  string `json:"token"`
	} `json:"data"`
}

type UpdateUserReq struct {
	Id     string `json:"id"` // 用户id
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Gender int64  `json:"gender"`
	Pic    string `json:"pic"`
}

type UpdateUserResp struct {
	Default
}

type User struct {
	Id         int64  `json:"id"` // 用户id
	Name       string `json:"name"`
	OpenId     string `json:"openid"`
	Phone      string `json:"phone"`
	Gender     int    `json:"gender"`
	IdNumber   string `json:"idNumber"`
	Pic        string `json:"pic"`
	CreateTime string `json:"createTime"`
}

type UserinfoReq struct {
	UserId string `path:"id"` // 用户id
}

type UserinfoResp struct {
	Default
	User struct {
		Id         int64  `json:"id"` // 用户id
		Name       string `json:"name"`
		OpenId     string `json:"openid"`
		Phone      string `json:"phone"`
		Gender     int    `json:"gender"`
		IdNumber   string `json:"idNumber"`
		Pic        string `json:"pic"`
		CreateTime string `json:"createTime"`
	} `json:"data"` // 用户信息
}
