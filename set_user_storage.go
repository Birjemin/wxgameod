package wxgameod

import (
	"fmt"
	"github.com/birjemin/wxgameod/utils"
	"log"
)

// SetUserStorage model
type SetUserStorage struct {
	AccessToken string
	SessionKey  string
	OpenID      string
	SigMethod   string
	KvList      string
	HTTPRequest *utils.HTTPClient
}

// RespSetUserStorage response
type RespSetUserStorage struct {
	CommonError
}

// SetUserStorage Set user storage
func (r *SetUserStorage) SetUserStorage() (*RespSetUserStorage, error) {
	return r.doSetUserStorage(wechatDomain)
}

// doSetUserStorage handle this action
func (r *SetUserStorage) doSetUserStorage(domain string) (*RespSetUserStorage, error) {

	url := fmt.Sprintf("%s%s?access_token=%s&signature=%s&openid=%s&sig_method=%s",
		domain,
		r.getSetUserStorageURI(),
		r.AccessToken,
		r.generateSignature(),
		r.OpenID,
		r.getSigMethod())

	var resp = new(RespSetUserStorage)

	if err := r.HTTPRequest.HTTPPostJSON(url, r.KvList); err != nil {
		log.Println("[set_user_storage]do, post failed", err)
		return resp, err
	}

	if err := r.HTTPRequest.GetResponseJSON(resp); err != nil {
		log.Println("[set_user_storage]do, response json failed", err)
		return resp, err
	}
	return resp, nil
}

// getSigMethod get sig method
func (r *SetUserStorage) getSigMethod() string {
	var sigMethod = "hmac_sha256"
	if r.SigMethod != "" {
		sigMethod = r.SigMethod
	}
	return sigMethod
}

// generateSignature cal signature
func (r *SetUserStorage) generateSignature() string {
	return utils.GenerateSha256(r.SessionKey, r.KvList)
}

// getInteractiveDataUrl get set_interactive_data URI
func (r *SetUserStorage) getSetUserStorageURI() string {
	return setUserStorageURI
}
