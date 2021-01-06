package wxgameod

import (
	"fmt"
	"github.com/birjemin/wxgameod/utils"
	"log"
)

// SetUserInteractiveData struct
type SetUserInteractiveData struct {
	AccessToken string
	SessionKey  string
	OpenID      string
	SigMethod   string
	KvList      string
	HTTPRequest *utils.HTTPClient
}

// RespSetUserInteractiveData response
type RespSetUserInteractiveData struct {
	CommonError
}

// SetUserInteractiveData set user interactive data
func (i *SetUserInteractiveData) SetUserInteractiveData() (*RespSetUserInteractiveData, error) {
	return i.doSetUserInteractiveData(wechatDomain)
}

// doSetUserInteractiveData do handle
func (i *SetUserInteractiveData) doSetUserInteractiveData(domain string) (*RespSetUserInteractiveData, error) {

	url := fmt.Sprintf("%s%s?access_token=%s&signature=%s&openid=%s&sig_method=%s",
		domain,
		i.getInteractiveDataURI(),
		i.AccessToken,
		i.generateSignature(),
		i.OpenID,
		i.getSigMethod())

	var resp = new(RespSetUserInteractiveData)

	if err := i.HTTPRequest.HTTPPostJSON(url, i.KvList); err != nil {
		log.Println("[set_user_interactive_data]do, post failed", err)
		return resp, err
	}

	// log.Println("url: ", url)
	// log.Println("params: ", i.InteractiveData)
	// log.Println("params: ", i.SessionKey)

	if err := i.HTTPRequest.GetResponseJSON(resp); err != nil {
		log.Println("[set_user_interactive_data]do, response json failed", err)
		return resp, err
	}
	return resp, nil
}

// getSigMethod get sig method
func (i *SetUserInteractiveData) getSigMethod() string {
	var sigMethod = "hmac_sha256"
	if i.SigMethod != "" {
		sigMethod = i.SigMethod
	}
	return sigMethod
}

// generateSignature cal signature
func (i *SetUserInteractiveData) generateSignature() string {
	return utils.GenerateSha256(i.SessionKey, i.KvList)
}

// getInteractiveDataUrl get set_interactive_data URI
func (i *SetUserInteractiveData) getInteractiveDataURI() string {
	return setUserInteractiveDataURI
}
