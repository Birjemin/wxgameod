package wxgameod

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	wechatDomain = "https://api.weixin.qq.com"

	removeUserStorageURI      = "/wxa/remove_user_storage"
	setUserInteractiveDataURI = "/wxa/setuserinteractivedata"
	setUserStorageURI         = "/wxa/set_user_storage"
)

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

// CommonError model
type CommonError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
